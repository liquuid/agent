package cli

import (
	"io/ioutil"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	"fmt"

	"github.com/subutai-io/agent/config"
	"github.com/subutai-io/agent/db"
	"github.com/subutai-io/agent/lib/fs"
	"github.com/subutai-io/agent/lib/gpg"
	ovs "github.com/subutai-io/agent/lib/net"
	"github.com/subutai-io/agent/log"
)

// MapPort exposes internal container ports to sockExt RH interface. It supports udp, tcp, http(s) protocols and other reverse proxy features
func MapPort(protocol, sockInt, sockExt, policy, domain, cert string, list, remove, sslbcknd bool) {
	if list {
		for _, v := range mapList(protocol) {
			fmt.Println(v)
		}
		return
	}

	if protocol != "tcp" && protocol != "udp" && protocol != "http" && protocol != "https" {
		log.Error("Unsupported protocol \"" + protocol + "\"")
	} else if protocol == "tcp" || protocol == "udp" {
		domain = protocol
	}

	if !ovs.ValidSocket(sockExt) {
		sockExt = "0.0.0.0:" + sockExt
	}

	switch {
	case (protocol == "http" || protocol == "https") && len(domain) == 0:
		log.Error("\"-d domain\" is mandatory for http protocol")
	case remove:
		mapRemove(protocol, sockExt, domain, sockInt)
	case protocol == "https" && (len(cert) == 0 || !gpg.ValidatePem(cert)):
		log.Error("\"-c certificate\" is missing or invalid pem file")
	case len(sockInt) != 0 && !ovs.ValidSocket(sockInt):
		log.Error("Invalid internal socket \"" + sockInt + "\"")
	case (strings.HasSuffix(sockExt, ":8443") || strings.HasSuffix(sockExt, ":8444") || strings.HasSuffix(sockExt, ":8086")) &&
		sockInt != "10.10.10.1:"+strings.Split(sockExt, ":")[1]:
		log.Error("Reserved system ports")
	case len(sockInt) != 0:
		// check sockExt port and create nginx config
		if portIsNew(protocol, sockInt, domain, &sockExt) {
			newConfig(protocol, sockExt, domain, cert, sslbcknd)
		}

		// add containers to backend
		addLine(config.Agent.DataPrefix+"nginx-includes/"+protocol+"/"+sockExt+"-"+domain+".conf",
			"#Add new host here", "	server "+sockInt+";", false)

		// save information to database
		saveMapToDB(protocol, sockExt, domain, sockInt)
		containerMapToDB(protocol, sockExt, domain, sockInt)
		balanceMethod(protocol, sockExt, domain, policy)

		if socket := strings.Split(sockExt, ":"); socket[0] == "0.0.0.0" {
			log.Info(ovs.GetIp() + ":" + socket[1])
		} else {
			log.Info(sockExt)
		}
	case len(policy) != 0:
		balanceMethod(protocol, sockExt, domain, policy)
	}
	restart()
}

func mapList(protocol string) (list []string) {
	bolt, err := db.New()
	log.Check(log.ErrorLevel, "Openning portmap database to get list", err)
	switch protocol {
	case "tcp", "udp", "http", "https":
		list = bolt.PortmapList(protocol)
	default:
		for _, v := range []string{"tcp", "udp", "http", "https"} {
			list = append(list, bolt.PortmapList(v)...)
		}
	}
	log.Check(log.WarnLevel, "Closing database", bolt.Close())
	return
}

func mapRemove(protocol, sockExt, domain, sockInt string) {
	bolt, err := db.New()
	log.Check(log.ErrorLevel, "Openning portmap database to remove mapping", err)
	defer bolt.Close()
	// Commenting this section out to insure config deletion even if db doesn't have it
	// if !bolt.PortInMap(protocol, sockExt, domain, sockInt) {
	// 	return
	// }
	log.Debug("Removing mapping: " + protocol + " " + sockExt + " " + domain + " " + sockInt)

	if bolt.PortMapDelete(protocol, sockExt, domain, sockInt) > 0 {
		if strings.Contains(sockInt, ":") {
			sockInt = sockInt + ";"
		} else {
			sockInt = sockInt + ":"
		}
		addLine(config.Agent.DataPrefix+"nginx-includes/"+protocol+"/"+sockExt+"-"+domain+".conf",
			"server "+sockInt, " ", true)
	} else {
		if bolt.PortMapDelete(protocol, sockExt, domain, "") == 0 {
			bolt.PortMapDelete(protocol, sockExt, "", "")
		}
		os.Remove(config.Agent.DataPrefix + "nginx-includes/" + protocol + "/" + sockExt + "-" + domain + ".conf")
		if protocol == "https" {
			os.Remove(config.Agent.DataPrefix + "web/ssl/https-" + sockExt + "-" + domain + ".key")
			os.Remove(config.Agent.DataPrefix + "web/ssl/https-" + sockExt + "-" + domain + ".crt")
		}
	}
}

func isFree(protocol, sockExt string) bool {
	switch protocol {
	case "tcp", "http", "https":
		if ln, err := net.Listen("tcp", sockExt); err == nil {
			ln.Close()
			return true
		}
	case "udp":
		if addr, err := net.ResolveUDPAddr("udp", sockExt); err == nil {
			if ln, err := net.ListenUDP("udp", addr); err == nil {
				ln.Close()
				return true
			}
		}
	}
	return false
}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func portIsNew(protocol, sockInt, domain string, sockExt *string) bool {
	socket := strings.Split((*sockExt), ":")
	if len(socket) > 1 && socket[1] != "" {
		if port, err := strconv.Atoi(socket[1]); err != nil || port < 1000 || port > 65536 {
			if !(strings.Contains(protocol, "http") && (port == 80 || port == 443)) {
				log.Error("Port number in \"external\" should be integer in range of 1000-65536")
			}
		}
		if isFree(protocol, (*sockExt)) {
			return true
		}

		bolt, err := db.New()
		defer bolt.Close()
		log.Check(log.ErrorLevel, "Opening portmap database to read existing mappings", err)
		if !bolt.PortInMap(protocol, (*sockExt), "", "") && socket[1] != "80" {
			log.Error("Port is busy")
		} else if bolt.PortInMap(protocol, (*sockExt), domain, sockInt) {
			log.Error("Map is already exists")
		}
		return !bolt.PortInMap(protocol, (*sockExt), domain, "")
	}
	for port := strconv.Itoa(random(1000, 65536)); isFree(protocol, socket[0]+":"+port); port = strconv.Itoa(random(1000, 65536)) {
		(*sockExt) = socket[0] + ":" + port
		return true
	}
	return false
}

func newConfig(protocol, sockExt, domain, cert string, sslbcknd bool) {
	log.Check(log.WarnLevel, "Creating nginx include folder",
		os.MkdirAll(config.Agent.DataPrefix+"nginx-includes/"+protocol, 0755))
	conf := config.Agent.DataPrefix + "nginx-includes/" + protocol + "/" + sockExt + "-" + domain + ".conf"

	switch protocol {
	case "https":
		log.Check(log.ErrorLevel, "Creating certificate dirs", os.MkdirAll(config.Agent.DataPrefix+"/web/ssl/", 0755))
		fs.Copy(config.Agent.AppPrefix+"etc/nginx/tmpl/vhost-ssl.example", conf)
		addLine(conf, "return 301 https://$host$request_uri;  # enforce https", "	    return 301 https://$host:"+strings.Split(sockExt, ":")[1]+"$request_uri;  # enforce https", true)
		addLine(conf, "listen	443;", "	listen "+sockExt+";", true)
		addLine(conf, "server_name DOMAIN;", "	server_name "+domain+";", true)
		if sslbcknd {
			addLine(conf, "proxy_pass http://DOMAIN-upstream/;", "	proxy_pass https://https-"+strings.Replace(sockExt, ":", "-", -1)+"-"+domain+";", true)
		} else {
			addLine(conf, "proxy_pass http://DOMAIN-upstream/;", "	proxy_pass http://https-"+strings.Replace(sockExt, ":", "-", -1)+"-"+domain+";", true)
		}
		addLine(conf, "upstream DOMAIN-upstream {", "upstream https-"+strings.Replace(sockExt, ":", "-", -1)+"-"+domain+" {", true)

		crt, key := gpg.ParsePem(cert)
		log.Check(log.WarnLevel, "Writing certificate body", ioutil.WriteFile(config.Agent.DataPrefix+"web/ssl/https-"+sockExt+"-"+domain+".crt", crt, 0644))
		log.Check(log.WarnLevel, "Writing key body", ioutil.WriteFile(config.Agent.DataPrefix+"web/ssl/https-"+sockExt+"-"+domain+".key", key, 0644))

		addLine(conf, "ssl_certificate /var/snap/subutai/current/web/ssl/UNIXDATE.crt;",
			"ssl_certificate "+config.Agent.DataPrefix+"web/ssl/https-"+sockExt+"-"+domain+".crt;", true)
		addLine(conf, "ssl_certificate_key /var/snap/subutai/current/web/ssl/UNIXDATE.key;",
			"ssl_certificate_key "+config.Agent.DataPrefix+"web/ssl/https-"+sockExt+"-"+domain+".key;", true)
	case "http":
		fs.Copy(config.Agent.AppPrefix+"etc/nginx/tmpl/vhost.example", conf)
		addLine(conf, "listen 	80;", "	listen "+sockExt+";", true)
		addLine(conf, "return 301 http://$host$request_uri;", "	    return 301 http://$host:"+strings.Split(sockExt, ":")[1]+"$request_uri;", true)
		addLine(conf, "server_name DOMAIN;", "	server_name "+domain+";", true)
		addLine(conf, "proxy_pass http://DOMAIN-upstream/;", "	proxy_pass http://http-"+strings.Replace(sockExt, ":", "-", -1)+"-"+domain+";", true)
		addLine(conf, "upstream DOMAIN-upstream {", "upstream http-"+strings.Replace(sockExt, ":", "-", -1)+"-"+domain+" {", true)
		if strings.HasSuffix(sockExt, ":80") {
			httpRedirect(sockExt, domain)
		}
	case "tcp":
		fs.Copy(config.Agent.AppPrefix+"etc/nginx/tmpl/stream.example", conf)
		addLine(conf, "listen PORT;", "	listen "+sockExt+";", true)
	case "udp":
		fs.Copy(config.Agent.AppPrefix+"etc/nginx/tmpl/stream.example", conf)
		addLine(conf, "listen PORT;", "	listen "+sockExt+" udp;", true)
	}
	addLine(conf, "server localhost:81;", " ", true)
	addLine(conf, "upstream PROTO-PORT {", "upstream "+protocol+"-"+strings.Replace(sockExt, ":", "-", -1)+"-"+domain+" {", true)
	addLine(conf, "proxy_pass PROTO-PORT;", "	proxy_pass "+protocol+"-"+strings.Replace(sockExt, ":", "-", -1)+"-"+domain+";", true)
}

func balanceMethod(protocol, sockExt, domain, policy string) {
	replaceString := "upstream " + protocol + "-" + strings.Replace(sockExt, ":", "-", -1) + "-" + domain + " {"
	replace := false
	bolt, err := db.New()
	log.Check(log.ErrorLevel, "Openning portmap database to check if port is mapped", err)
	if !bolt.PortInMap(protocol, sockExt, domain, "") {
		log.Error("Port is not mapped")
	}
	switch policy {
	case "round-robin", "round_robin":
		policy = "#round-robin"
	//  "least_conn":
	case "least_time":
		if protocol == "tcp" {
			policy = policy + " connect"
		} else {
			policy = policy + " header"
			log.Warn("This policy is not supported in http upstream")
			return
		}
	case "hash":
		policy = policy + " $remote_addr"
	case "ip_hash":
		if protocol != "http" {
			log.Warn("ip_hash policy allowed only for http protocol")
			return
		}
	default:
		log.Debug("Unsupported balancing method \"" + policy + "\", ignoring")
		return
	}

	if p := bolt.GetMapMethod(protocol, sockExt, domain); len(p) != 0 && p != policy {
		replaceString = "; #policy"
		replace = true
	} else if p == policy {
		return
	}
	log.Check(log.WarnLevel, "Saving map method", bolt.SetMapMethod(protocol, sockExt, domain, policy))
	log.Check(log.WarnLevel, "Closing database", bolt.Close())

	addLine(config.Agent.DataPrefix+"nginx-includes/"+protocol+"/"+sockExt+"-"+domain+".conf",
		replaceString, "	"+policy+"; #policy", replace)
}

func httpRedirect(sockExt, domain string) {
	var redirect = `server {
	    listen      80; #redirect
    	server_name ` + domain + `;
    	return 301 http://$host$request_uri;
}`

	addLine(config.Agent.DataPrefix+"nginx-includes/http/"+sockExt+"-"+domain+".conf",
		"#redirect placeholder", redirect, true)

}

func saveMapToDB(protocol, sockExt, domain, sockInt string) {
	bolt, err := db.New()
	log.Check(log.ErrorLevel, "Openning database to save portmap", err)
	if !bolt.PortInMap(protocol, sockExt, domain, sockInt) {
		log.Check(log.WarnLevel, "Saving port map to database", bolt.PortMapSet(protocol, sockExt, domain, sockInt))
	}
	log.Check(log.WarnLevel, "Closing database", bolt.Close())
}

func containerMapToDB(protocol, sockExt, domain, sockInt string) {
	bolt, err := db.New()
	log.Check(log.ErrorLevel, "Openning database to add portmap to container", err)
	for _, name := range bolt.ContainerByKey("ip", strings.Split(sockInt, ":")[0]) {
		bolt.ContainerMapping(name, protocol, sockExt, domain, sockInt)
	}
	log.Check(log.WarnLevel, "Closing database", bolt.Close())
}
