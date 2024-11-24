# HEAD-to-GET Plugin

This plugin converts any HTTP request that has the `HEAD` verb/method to `GET`, useful for certain Java or .Net web apps that don't support `HEAD` but you use an external health check that wants to send HEAD anyways.

## Configuration

1. Enable the plugin in your Traefik configuration:

```yaml
experimental:
 plugins:
  head_to_get:
   moduleName: github.com/cdwiegand/traefik-head-to-get
   version: v0.2.1
```

1. Define the middleware. Note that this plugin does not need any configuration, however, values must be passed in for it to be accepted within Traefik:

```yaml
http:
 # ...
 middlewares:
  # this name must match the middleware that you attach to routers later
  head-to-get:
    plugin:
      head_to_get:
        removeExposingHeaders: true
```

Please note that traefik requires at least one configuration variable set, to keep the defaults you can set `addCspHeader: true` to accomodate this. *This is not a requirement of this plugin, but a traefik requirement.*

1. Then add it to your given routers, such as this:

```yaml
http:
 # ...
 routers:
  example-router:
   rule: host(`demo.localhost`)
   service: service-foo
   entryPoints:
    - web
   # add these 2 lines, use the same name you defined directly under "middlewares":
   middlewares: 
    - head-to-get
   # end add those 2 lines
```

1. You are done!

## Testing Methods

Testing by using local plugin functionality, assuming the code is checked out to `C:\devel\traefik-head-to-get`:

```bash
docker run --rm -it -p 8888:80 -v C:\devel\traefik-head-to-get\:/srv/plugins-local/src/github.com/cdwiegand/traefik-head-to-get:ro -w /srv traefik:3.0 --entryPoints.web.address=:80 --experimental.localPlugins.head_to_get.modulename=github.com/cdwiegand/traefik-head-to-get --providers.file.filename=/srv/plugins-local/src/github.com/cdwiegand/traefik-head-to-get/testing.traefik.yml --api=true --api.dashboard=true
```

and go to <http://localhost:8888/dashboard/> and inspect the browser's Network tab to see the Server header in the response replaced with "Nope/2.0".
