Hub API
=======
Subutai Agent private REST API

**Version:** 1.0.0

### /rh/id
---
##### ***GET***
**Description:** Returns JSON formatted Id of RH, UUID which is the PGP fingerprint

**Responses**

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | JSON formatted Id of RH | [ [Fingerprint](#fingerprint) ] |
| default | generic error response |  |

### /list
---
##### ***GET***
**Description:** Info returns JSON formatted list of Subutai instances with information such as IP address, parent template, etc.

**Parameters**

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| limit | query | Limit of returned entries on response | No | integer |
| containersOnly | query | Parameter to return containers only | No | boolean |
| templatesOnly | query | Parameter to return templates only | No | boolean |
| withAncestors | query | Parameter to return containers/templates with ancestors | No | boolean |
| withParents | query | Parameter to return containers/templates with parents | No | boolean |

**Responses**

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | List all subutai instances | [ [Container](#container) ] |
| default | generic error response |  |

### /attach/{container}
---
##### ***GET***
**Description:** LxcAttach allows user to use container's TTY.
`name` should be available running Subutai container, otherwise command will return error message and non-zero exit code.

**Parameters**

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| container | path | Container to be attached | Yes | string |

**Responses**

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Ok | [ [Message](#message) ] |
| default | generic error response |  |

### /container/{name}
---
##### ***GET***
**Description:** Get container Info in JSON formatted object

**Parameters**

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| name | path | Name of container | Yes | string |
| containersOnly | query | Parameter to return containers only | No | boolean |
| templatesOnly | query | Parameter to return templates only | No | boolean |
| withAncestors | query | Parameter to return containers/templates with ancestors | No | boolean |
| withParents | query | Parameter to return containers/templates with parents | No | boolean |
| detailedInfo | query | Parameter to return detailed info | No | boolean |

**Responses**

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [Container](#container) |
| default | error | [Error](#error) |

##### ***DELETE***
**Parameters**

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| name | path | Name of container | Yes | string |
| containersOnly | query | Parameter to return containers only | No | boolean |
| templatesOnly | query | Parameter to return templates only | No | boolean |
| withAncestors | query | Parameter to return containers/templates with ancestors | No | boolean |
| withParents | query | Parameter to return containers/templates with parents | No | boolean |
| detailedInfo | query | Parameter to return detailed info | No | boolean |

**Responses**

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 204 | Deleted |
| default | error | [Error](#error) |

### /backup/{name}
---
##### ***GET***
**Description:** BackupContainer takes a snapshots of each container's volume and stores it in the `/mnt/backups/container_name/datetime/` directory. A full backup creates a delta-file of each BTRFS subvolume. An incremental backup (default) creates a delta-file with the difference of changes between the current and last snapshots. All deltas are compressed to archives in `/mnt/backups/` directory (container_datetime.tar.gz or container_datetime_Full.tar.gz for full backup). A changelog file can be found next to backups archive (container_datetime_changelog.txt or container_datetime_Full_changelog.txt) which contains a list of changes made between two backups.

**Parameters**

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| name | path |  | Yes | string |
| full | query |  | No | boolean |
| stop | query |  | No | boolean |

**Responses**

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [Message](#message) |
| default | error | [Error](#error) |

### /batch
---
##### ***POST***
**Description:** Batch binding provides a mechanism to perform several Subutai commands in the container in batch, passed in a single JSON message. Initially, the purpose of this command was internal for SS <-> Agent communication, yet it may be invoked manually from the CLI. The response from a batch command returns a JSON array with each element representing the results (response) from each command (request) in the batch: the positions of responses correlate with the request position in the array

**Parameters**

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| body | body |  | No | [Batchline](#batchline) |

**Responses**

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [Message](#message) |
| default | error | [Error](#error) |

### /cleanup
---
##### ***GET***
**Description:** Cleanup simply removes every resource associated with a Subutai container or template: data, network, configs, etc. The destroy command always runs each step in "force" mode to provide reliable deletion results; even if some instance components were already removed, the destroy command will continue to perform all operations once again while ignoring possible underlying errors: i.e. missing configuration files.

**Responses**

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [Message](#message) |
| default | error | [Error](#error) |

### /clone/{parent}/{child}
---
##### ***POST***
**Description:** Clone function creates new `child` container from a Subutai `parent` template. If the specified template argument is not deployed in system, Subutai first tries to import it, and if import succeeds, it then continues to clone from the imported template image. By default, clone will use the NAT-ed network interface with IP address received from the Subutai DHCP server, but this behavior can be changed with command options described below.
If `ipaddr` option is defined, separate bridge interface will be created in specified VLAN and new container will receive static IP address. Option `envID` writes the environment ID string inside new container. Option `token` is intended to check the origin of new container creation request during environment build. This is one of the security checks which makes sure that each container creation request is authorized by registered user.
Option `kurjunToken` set kurjun token to clone private and shared templates
The clone options are not intended for manual use: unless you're confident about what you're doing. Use default clone format without additional options to create Subutai containers.

**Parameters**

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| parent | path | Parent name | Yes | string |
| child | path | Child name | Yes | string |
| body | body |  | No | [Cloneargs](#cloneargs) |

**Responses**

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [Message](#message) |
| default | error | [Error](#error) |

### /config
---
##### ***POST***
**Description:** Allows read and write container's configuration file

**Parameters**

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| body | body |  | No | [Configoptions](#configoptions) |

**Responses**

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [Message](#message) |
| default | error | [Error](#error) |

##### ***DELETE***
**Description:** Delete entry in configuration file

**Parameters**

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| body | body |  | No | [Configoptions](#configoptions) |

**Responses**

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 204 | Deleted |
| default | error | [Error](#error) |

### /demote/{container}
---
##### ***POST***
**Description:** Converts template into regular Subutai container.
A Subutai template is a "locked down" container only to be used for cloning purposes. It cannot be started, and its file system cannot be modified: it's read-only. Normal operational containers are promoted into templates, but sometimes you might want to demote them back to regular containers. This is what the demote sub command does: it reverts a template without children back into a normal container. Demoted container will use NAT network interface and dynamic IP address if opposite options are not specified.

**Parameters**

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| container | path | Container name | Yes | string |
| body | body |  | No | [Demoteoptions](#demoteoptions) |

**Responses**

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [Message](#message) |
| default | error | [Error](#error) |

### /promote/{container}
---
##### ***POST***
**Description:** Promote turns a Subutai container into container template which may be cloned with "clone" command. Promote executes several simple steps, such as dropping a container's configuration to default values, dumping the list of installed packages (this step requires the target container to still be running), and setting the container's filesystem to read-only to prevent changes.

**Parameters**

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| container | path | Container name | Yes | string |
| body | body |  | No | [Promoteoptions](#promoteoptions) |

**Responses**

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [Message](#message) |
| default | error | [Error](#error) |

### /export/{container}
---
##### ***POST***
**Description:** Export prepares an archive from a template in the `/mnt/lib/lxc/tmpdir/` path. This archive can be moved to another Subutai peer and deployed as ready-to-use template or uploaded to Subutai's global template repository to make it widely available for others to use.
Export consist of two steps if the target is a container: container promotion to template (see "promote" command) and packing the template into the archive. If already a template just the packing of the archive takes place.
Configuration values for template metadata parameters can be overridden on export, like the recommended container size when the template is cloned using `-s` option.

**Parameters**

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| container | path | Container name | Yes | string |
| body | body |  | No | [Exportoptions](#exportoptions) |

**Responses**

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [Message](#message) |
| default | error | [Error](#error) |

### /hostname/{container}/{name}
---
##### ***POST***
**Description:**         - type: string name: container in: path required: true

**Parameters**

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| container | path | Container name | Yes | string |
| name | path | New name | Yes | string |

**Responses**

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [Message](#message) |
| default | error | [Error](#error) |

### /destroy/{ID}
---
##### ***GET***
**Description:** Destroy Subutai container

**Parameters**

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| ID | path | container name or vlan id | Yes | string |
| vlan | query | If true the ID is a vlan, if false is a container name | No | boolean |

**Responses**

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Ok | [ [Message](#message) ] |
| default | generic error response |  |

### /import/{container}
---
##### ***GET***
**Description:** Import Subutai template

**Parameters**

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| container | path | Container name | Yes | string |
| body | body |  | No | [Importoptions](#importoptions) |

**Responses**

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Ok | [ [Message](#message) ] |
| default | generic error response |  |

### /info
---
##### ***GET***
**Description:** Info command's purposed is to display common system information, such as external IP address to access the container host quotas, its CPU model, RAM size, etc. It's mainly used for internal SS needs.

**Parameters**

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| command | query | Commands available:  ipaddr, ports, os, quota, system | Yes | string |

**Responses**

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Ok | [ [InfoHostStat](#infohoststat) ] |
| default | generic error response |  |

### /metrics
---
##### ***GET***
**Description:** HostMetrics function retrieves monitoring data from a time-series database deployed in the SS Management server for container hosts and Resource Hosts. Statistics are being collected by the Subutai daemon and includes common information like CPU utilization, network load, RAM and disk usage for both containers and hosts. Since the database is located on the SS Management Host, hosts which are not a part of a Subutai peer have no access to this information. Data aggregation in the time-series database has following configuration: last hour statistic is stored "as is", last day data aggregates to 1 minute interval, last week is in 5 minute intervals, After 7 days all statistics is are overwritten by new incoming data.

**Parameters**

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| start | query | Start date. Date Example: 2006-01-02 15:04:05 | Yes | string |
| end | query | End date.  Date Example: 2006-01-02 15:04:05 | Yes | string |

**Responses**

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Ok | [ [MetricInfo](#metricinfo) ] |
| default | generic error response |  |

### /quota
---
##### ***GET***
**Description:** Quota function controls container's quotas and thresholds. Available resources: cpu, % cpuset, available cores ram, Mb network, Kbps rootfs/home/var/opt, Gb The threshold value represents a percentage for each resource. Once resource consumption exceeds this threshold it triggers an alert. The clone operation, sets no quotas and thresholds for new containers; quotas need to be configured with quota command after a clone operation.

**Parameters**

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| body | body |  | No | [quotaArgs](#quotaargs) |

**Responses**

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Ok | [ [quotaOutput](#quotaoutput) ] |
| default | generic error response |  |

### /rename/{source}/{newname}
---
##### ***GET***
**Description:** Renames a Subutai container impacting filesystem paths, configuration values, etc.

**Parameters**

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| source | path | Source name | Yes | string |
| newname | path | New name | Yes | string |

**Responses**

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Ok | [ [Message](#message) ] |
| default | generic error response |  |

### /restore/{container}
---
##### ***GET***
**Description:** RestoreContainer restores a Subutai container to a snapshot at a specified timestamp if such a backup archive is available.

**Parameters**

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| container | path | Container name | Yes | string |
| date | query | Date argument | No | string |
| newcontainer | query | Name of new container | No | string |

**Responses**

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Ok | [ [Message](#message) ] |
| default | generic error response |  |

### /start/{container}
---
##### ***GET***
**Description:** Starts a Subutai container and checks if container state changed to "running" or "starting". If state is not changing for 60 seconds, then the "start" operation is considered to have failed.

**Parameters**

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| container | path | Container name | Yes | string |

**Responses**

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Ok | [ [Message](#message) ] |
| default | generic error response |  |

### /stop/{container}
---
##### ***GET***
**Description:** Stops a Subutai container with an additional state check.

**Parameters**

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| container | path | Container name | Yes | string |

**Responses**

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Ok | [ [Message](#message) ] |
| default | generic error response |  |

### /update/{container}
---
##### ***GET***
**Description:** Update operation can be divided into two different types: container updates and Resource Host updates. Container updates simply perform apt-get update and upgrade operations inside target containers without any extra commands. Since SS Management is just another container, the Subutai update command works fine with the management container too.
The second type of update, a Resource Host update, checks the Ubuntu Store and compares available snap packages with those currently installed in the system and, if a newer version is found, installs it. Please note, system security policies requires that such commands should be performed by the superuser manually, otherwise an application's attempt to update itself will be blocked.

**Parameters**

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| container | path | Container name | Yes | string |
| check | query | Just check for updates, don't update | Yes | boolean |

**Responses**

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Ok | [ [Message](#message) ] |
| default | generic error response |  |

### /p2p
---
##### ***GET***
**Description:** P2P function controls and configures the peer-to-peer network structure: the swarm which includes all hosts with same the same swarm hash and secret key.
P2P is a base layer for Subutai environment networking: all containers in same environment are connected to each other via VXLAN tunnels and are accesses as if they were in one LAN. It doesn't matter where the containers are physically located.

**Parameters**

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| command | query | Accepted commands: - list: list of p2p instances - interfaces: list of p2p interfaces - peers: list of p2p swarm participants by hash - version: print p2p version | Yes | string |

**Responses**

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Ok | [ [Text](#text) ] |
| default | generic error response |  |

##### ***PUT***
**Description:** P2P function controls and configures the peer-to-peer network structure: the swarm which includes all hosts with same the same swarm hash and secret key.
P2P is a base layer for Subutai environment networking: all containers in same environment are connected to each other via VXLAN tunnels and are accesses as if they were in one LAN. It doesn't matter where the containers are physically located.

**Parameters**

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| body | body |  | No | [P2pArgs](#p2pargs) |

**Responses**

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Ok | [ [Text](#text) ] |
| default | generic error response |  |

##### ***DELETE***
**Description:** P2P function controls and configures the peer-to-peer network structure: the swarm which includes all hosts with same the same swarm hash and secret key.
P2P is a base layer for Subutai environment networking: all containers in same environment are connected to each other via VXLAN tunnels and are accesses as if they were in one LAN. It doesn't matter where the containers are physically located.

**Parameters**

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| body | body |  | No | [P2pArgs](#p2pargs) |

**Responses**

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 204 | Deleted |
| default | generic error response |  |

##### ***POST***
**Description:** P2P function controls and configures the peer-to-peer network structure: the swarm which includes all hosts with same the same swarm hash and secret key.
P2P is a base layer for Subutai environment networking: all containers in same environment are connected to each other via VXLAN tunnels and are accesses as if they were in one LAN. It doesn't matter where the containers are physically located.

**Parameters**

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| body | body |  | No | [P2pArgs](#p2pargs) |

**Responses**

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 201 | Created | [Item](#item) |
| default | generic error response |  |

### /proxy
---
##### ***GET***
**Description:** ProxyCheck exits with 0 code if domain or node is exists in specified vlan, otherwise exitcode is 1

**Parameters**

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| body | body |  | No | [ProxyArgs](#proxyargs) |

**Responses**

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Ok | [ [Message](#message) ] |
| default | generic error response |  |

##### ***DELETE***
**Description:** ProxyDel checks what need to be removed - domain or node and pass args to required functions

**Parameters**

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| body | body |  | No | [ProxyArgs](#proxyargs) |

**Responses**

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 204 | Deleted |
| default | generic error response |  |

##### ***POST***
**Description:** ProxyAdd checks input args and perform required operations to configure reverse proxy

**Parameters**

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| body | body |  | No | [ProxyArgs](#proxyargs) |

**Responses**

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 201 | Created | [Message](#message) |
| default | generic error response |  |

### /tunnel/list
---
##### ***GET***
**Description:** performs tunnel check and shows "alive" tunnels

**Responses**

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Ok | [ [Item](#item) ] |
| default | generic error response |  |

### /tunnel/check
---
##### ***GET***
**Description:** reads list, checks tunnel ttl, its state and then adds or removes required tunnels

**Responses**

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Ok | [ [Item](#item) ] |
| default | generic error response |  |

### /tunnel/{socket}
---
##### ***DELETE***
**Description:** TunDel removes tunnel entry from list and kills running tunnel process

**Parameters**

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| socket | path | Socket name | Yes | string |
| pid | query | PID | No | string |

**Responses**

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 204 | Deleted |
| default | generic error response |  |

##### ***POST***
**Description:** TunAdd adds tunnel to specified network socket

**Parameters**

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| socket | path | Socket name | Yes | string |
| timeout | query | Tunnels may also be set to be permanent (default) or temporary (ttl in seconds). The default destination port is 22. | No | string |
| global | query | There are two types of channels - local (default), which is created from destination address to host and global, from destination to Subutai Helper node. | No | boolean |

**Responses**

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 201 | Created | [Message](#message) |
| default | generic error response |  |

### /vxlan/list
---
##### ***GET***
**Description:** prints a list of existing VXLAN tunnels

**Responses**

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Ok | [ [Item](#item) ] |
| default | generic error response |  |

### /vxlan/{iface}
---
##### ***DELETE***
**Description:** removes OVS bridges and ports by name, brings system interface down

**Parameters**

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| iface | path | Network interface name | Yes | string |

**Responses**

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 204 | Deleted |
| default | generic error response |  |

### /vxlan/{tunnel}
---
##### ***POST***
**Description:** Creates VXLAN tunnel

**Parameters**

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| tunnel | path | Tunnel name | Yes | string |
| remoteip | query | Remote IP | Yes | string |
| vlan | query | Virtual local network | Yes | string |
| vni | query | Virtual Network Interface | Yes | string |

**Responses**

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 201 | Created | [Message](#message) |
| default | generic error response |  |

### Models
---

### Error  

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| code | long | Error code | No |
| message | string | Error message | Yes |

### Fingerprint  

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| hash | string | fingerprint | Yes |

### Container  

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| name | string | Name of container | Yes |
| parent | string | Name of container's parent | No |
| ancestor | string | Name of container's ancestor | No |

### Message  

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| text | string | Message returned | Yes |
| exitcode | string | Returned Exit code | No |

### Batchline  

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| action | string | Action to be executed list, backup, import, promote etc. | Yes |
| args | [ [Item](#item) ] | Arguments to action command | Yes |

### Item  

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| text | string |  | Yes |

### Cloneargs  

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| parent | string | Container/Template parent name | Yes |
| child | string | Child name | Yes |
| envID | string | Environment ID | No |
| ipaddr | string | IP address | No |
| token | string |  | No |
| kurjunToken | string | Kurjun Token | No |

### Configoptions  

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| key | string | Key | Yes |
| value | string | Value | Yes |

### Demoteoptions  

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| ipaddr | string | Ip address | No |
| vlan | string | VLAN name | No |

### Promoteoptions  

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| source | string | Source name | No |

### Exportoptions  

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| version | string | Version number | No |
| size | string | Size | No |
| token | string | Token | No |
| description | string | Description | No |
| private | boolean | Export to private template | No |

### Importoptions  

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| version | string | Version | No |
| torrent | boolean | Get template using torrent | No |
| token | string | Token | No |

### InfoHostStat  

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| cpu | [ [cpuInfo](#cpuinfo) ] |  | No |
| disk | [ [diskInfo](#diskinfo) ] |  | No |
| ram | [ [ramInfo](#raminfo) ] |  | No |
| quota | [ [quotaInfo](#quotainfo) ] |  | No |
| hostname | string |  | No |

### MetricInfo  

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| series | string |  | No |
| messages | string |  | No |

### cpuInfo  

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| model | string | CPU model | No |
| coreCount | string | Number of CPU Cores | No |
| idle | string | CPU Idle | No |
| frequency | string | CPU frequency | No |

### diskInfo  

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| total | string | Total space available | No |
| used | string | Used space | No |

### ramInfo  

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| free | string |  | No |
| total | string |  | No |
| cached | string |  | No |

### quotaInfo  

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| container | string |  | No |
| cpu | string |  | No |
| ram | string |  | No |
| disk | [ [diskStructInfo](#diskstructinfo) ] |  | No |

### diskStructInfo  

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| resource | string | Resources available | No |

### quotaArgs  

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| name | string |  | Yes |
| resource | string |  | Yes |
| size | string |  | No |
| threshold | string |  | No |

### quotaOutput  

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| quota | string |  | Yes |
| threshold | string |  | Yes |

### P2pArgs  

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| interfaceName | string | Interface Name | Yes |
| hash | string | Hash | Yes |
| key | string | Key | Yes |
| ttl | string | Life time | Yes |
| localPeepIPAddr | string | Local Peep IP Address | No |
| portRange | string | Port Range | No |

### Text  

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| text | string |  | No |

### ProxyArgs  

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| vlan | string | VLan Name | Yes |
| node | string | Node Name | No |
| policy | string |  | No |
| cert | string | SSL certificate | No |
| domain | string | Domain name | No |