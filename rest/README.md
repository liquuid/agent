### How to generate swagger.yml
Install the node tool:

```bash
  npm install -g multi-file-swagger
```

Run the command like so:

```bash
  multi-file-swagger -o yaml index.yml > compiled.yml
```
Which can then be used for code generation and so on.