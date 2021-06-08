package container

import (
    "log"

    endure "github.com/spiral/endure/pkg/container"
    // plugins
    // "github.com/spiral/roadrunner-binary/v2/cli"
    "github.com/outway/roadrunner-binary/v2/internal/cli"
    httpPlugin "github.com/spiral/roadrunner/v2/plugins/http"
    "github.com/spiral/roadrunner/v2/plugins/informer"
    //"github.com/spiral/roadrunner/v2/plugins/kv/boltdb"
    //"github.com/spiral/roadrunner/v2/plugins/kv/memcached"
    //"github.com/spiral/roadrunner/v2/plugins/kv/memory"
    "github.com/spiral/roadrunner/v2/plugins/logger"
    "github.com/spiral/roadrunner/v2/plugins/metrics"
    "github.com/spiral/roadrunner/v2/plugins/redis"
    "github.com/spiral/roadrunner/v2/plugins/reload"
    "github.com/spiral/roadrunner/v2/plugins/resetter"
    "github.com/spiral/roadrunner/v2/plugins/rpc"
    "github.com/spiral/roadrunner/v2/plugins/server"
    "github.com/temporalio/roadrunner-temporal/activity"
    temporalClient "github.com/temporalio/roadrunner-temporal/client"
    "github.com/temporalio/roadrunner-temporal/workflow"
)

// Plugins returns active plugins for the endure container. Feel free to add or remove any plugins.
func Plugins() []interface{} {
	return []interface{}{
		// logger plugin
		&logger.ZapLogger{},
		// metrics plugin
		&metrics.Plugin{},
		// http server plugin
		&httpPlugin.Plugin{},
		// reload plugin
		&reload.Plugin{},
		// informer plugin (./rr workers, ./rr workers -i)
		&informer.Plugin{},
		// resetter plugin (./rr reset)
		&resetter.Plugin{},
		// rpc plugin (workers, reset)
		&rpcPlugin.Plugin{},
		// server plugin (NewWorker, NewWorkerPool)
		&server.Plugin{},
		// service plugin
		&service.Plugin{},

		// websockets broadcast bundle
		&websockets.Plugin{},
		&redis.Plugin{},
		&memory.Plugin{},
		// =========

		// plugin to serve static files
		&static.Plugin{},
		// headers
		&headers.Plugin{},
		// checker
		&status.Plugin{},
		// gzip
		&gzip.Plugin{},

		// temporal plugins
		&activity.Plugin{},
		&workflow.Plugin{},
		&temporalClient.Plugin{},
	}
}
