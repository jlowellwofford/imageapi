// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http",
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Mange system image containers",
    "title": "Image API",
    "version": "1.0.0"
  },
  "basePath": "/imageapi/v1",
  "paths": {
    "/attach/rbd": {
      "get": {
        "tags": [
          "attach"
        ],
        "operationId": "list_rbds",
        "responses": {
          "200": {
            "description": "list of rbd maps",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/rbd"
              }
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "post": {
        "tags": [
          "attach"
        ],
        "operationId": "map_rbd",
        "parameters": [
          {
            "name": "rbd",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/rbd"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "RBD attach succeed",
            "schema": {
              "$ref": "#/definitions/rbd"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/attach/rbd/{id}": {
      "get": {
        "tags": [
          "attach"
        ],
        "operationId": "get_rbd",
        "responses": {
          "200": {
            "description": "RBD entry",
            "schema": {
              "$ref": "#/definitions/rbd"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "delete": {
        "tags": [
          "attach"
        ],
        "operationId": "unmap_rbd",
        "responses": {
          "204": {
            "description": "Unmapped"
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "integer",
          "format": "int64",
          "name": "id",
          "in": "path",
          "required": true
        }
      ]
    },
    "/container": {
      "get": {
        "description": "Get a list of containers",
        "tags": [
          "containers"
        ],
        "operationId": "list_containers",
        "responses": {
          "200": {
            "description": "List of containers",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/container"
              }
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "post": {
        "description": "Create a container",
        "tags": [
          "containers"
        ],
        "operationId": "create_container",
        "parameters": [
          {
            "name": "container",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/container"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Container creation succeed",
            "schema": {
              "$ref": "#/definitions/container"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/container/{id}": {
      "get": {
        "description": "Get a container definition",
        "tags": [
          "containers"
        ],
        "operationId": "get_container",
        "responses": {
          "200": {
            "description": "Container entry",
            "schema": {
              "$ref": "#/definitions/container"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "delete": {
        "description": "Delete a container defition.  \nThis will stop running containers.\n",
        "tags": [
          "containers"
        ],
        "operationId": "delete_container",
        "responses": {
          "204": {
            "description": "Container deleted"
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "integer",
          "format": "int64",
          "name": "id",
          "in": "path",
          "required": true
        }
      ]
    },
    "/container/{id}/{state}": {
      "get": {
        "description": "Request a (valid) state for a container. \nValid states to request include: ` + "`" + `running` + "`" + `, ` + "`" + `exited` + "`" + `, ` + "`" + `paused` + "`" + ` (paused is not yet implemented)\n",
        "tags": [
          "containers"
        ],
        "operationId": "set_container_state",
        "responses": {
          "200": {
            "description": "Container state changed",
            "schema": {
              "$ref": "#/definitions/container"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "integer",
          "format": "int64",
          "name": "id",
          "in": "path",
          "required": true
        },
        {
          "enum": [
            "running",
            "exited"
          ],
          "type": "string",
          "name": "state",
          "in": "path",
          "required": true
        }
      ]
    },
    "/mount/overlay": {
      "get": {
        "tags": [
          "mounts"
        ],
        "operationId": "list_mounts_overlay",
        "responses": {
          "200": {
            "description": "list of overlay mounts",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/mount_overlay"
              }
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "post": {
        "tags": [
          "mounts"
        ],
        "operationId": "mount_overlay",
        "parameters": [
          {
            "name": "mount",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/mount_overlay"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Overlay mount succeed",
            "schema": {
              "$ref": "#/definitions/mount_overlay"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/mount/overlay/{id}": {
      "get": {
        "tags": [
          "mounts"
        ],
        "operationId": "get_mount_overlay",
        "responses": {
          "200": {
            "description": "Overlay mount entry",
            "schema": {
              "$ref": "#/definitions/mount_overlay"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "delete": {
        "tags": [
          "mounts"
        ],
        "operationId": "unmount_overlay",
        "responses": {
          "204": {
            "description": "Unmounted"
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "integer",
          "format": "int64",
          "name": "id",
          "in": "path",
          "required": true
        }
      ]
    },
    "/mount/rbd": {
      "get": {
        "tags": [
          "mounts"
        ],
        "operationId": "list_mounts_rbd",
        "responses": {
          "200": {
            "description": "list of rbd mounts",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/mount_rbd"
              }
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "post": {
        "tags": [
          "mounts"
        ],
        "operationId": "mount_rbd",
        "parameters": [
          {
            "name": "mount",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/mount_rbd"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "RBD mount succeed",
            "schema": {
              "$ref": "#/definitions/mount_rbd"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/mount/rbd/{id}": {
      "get": {
        "tags": [
          "mounts"
        ],
        "operationId": "get_mount_rbd",
        "responses": {
          "200": {
            "description": "RBD mount entry",
            "schema": {
              "$ref": "#/definitions/mount_rbd"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "delete": {
        "tags": [
          "mounts"
        ],
        "operationId": "unmount_rbd",
        "responses": {
          "204": {
            "description": "Unmounted"
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "integer",
          "format": "int64",
          "name": "id",
          "in": "path",
          "required": true
        }
      ]
    }
  },
  "definitions": {
    "container": {
      "description": "The ` + "`" + `container` + "`" + ` option describes a minimally namespaced container.",
      "type": "object",
      "required": [
        "mount",
        "command"
      ],
      "properties": {
        "command": {
          "type": "string"
        },
        "id": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        },
        "logfile": {
          "type": "string",
          "readOnly": true
        },
        "mount": {
          "$ref": "#/definitions/mount"
        },
        "namespaces": {
          "description": "A list of Linux namespaces to use.\n\nNote: This is currently unused.  All containers currently get ` + "`" + `mnt` + "`" + ` and ` + "`" + `pid` + "`" + `.\n      It's here as a placeholder for future use.\n",
          "type": "array",
          "items": {
            "$ref": "#/definitions/container_namespace"
          }
        },
        "state": {
          "description": "When read, this contains the current container state. On creation, this requests the initial state (valid options: ` + "`" + `created` + "`" + ` or ` + "`" + `running` + "`" + `). The default is ` + "`" + `created` + "`" + `.\n",
          "$ref": "#/definitions/container_state"
        },
        "systemd": {
          "description": "When ` + "`" + `systemd` + "`" + ` is set to ` + "`" + `true` + "`" + `, we will assume that this container will run ` + "`" + `systemd` + "`" + `, and perform the necessary magic dance to make systemd run inside of the container. The default is ` + "`" + `false` + "`" + `.\n",
          "type": "boolean"
        }
      }
    },
    "container_namespace": {
      "description": "Linux namespace",
      "type": "string",
      "enum": [
        "cgroup",
        "ipc",
        "net",
        "mnt",
        "pid",
        "time",
        "user",
        "uts"
      ]
    },
    "container_state": {
      "description": "Valid container states",
      "type": "string",
      "enum": [
        "created",
        "running",
        "stopping",
        "exited",
        "dead"
      ]
    },
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "mount": {
      "description": "Generically address mounts by kind and ID",
      "type": "object",
      "required": [
        "kind",
        "id"
      ],
      "properties": {
        "id": {
          "type": "integer",
          "format": "int64"
        },
        "kind": {
          "type": "string",
          "enum": [
            "overlay",
            "rbd"
          ]
        }
      }
    },
    "mount_overlay": {
      "description": "` + "`" + `mount_overlay` + "`" + ` describes an Overlayfs mount.  All mount points must be RBD ID's.\nAt very least, ` + "`" + `lower` + "`" + ` must be specified.  If ` + "`" + `upper` + "`" + ` length is zero, no ` + "`" + `upper` + "`" + `\nmounts will be used.  ` + "`" + `workdir` + "`" + ` will be assigned automatically.\n\nOverlay mounts are identified by their ` + "`" + `lower` + "`" + ` ID.\n",
      "type": "object",
      "required": [
        "lower"
      ],
      "properties": {
        "id": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        },
        "lower": {
          "description": "This is an array of RBD IDs, interpreted in order for multiple lower dirs. At least one must be specified.",
          "type": "array",
          "items": {
            "type": "integer",
            "format": "int64"
          }
        },
        "mountpoint": {
          "type": "string",
          "readOnly": true
        },
        "ref": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        },
        "upperdir": {
          "description": "currently, upperdir is always a directory in mountDir",
          "type": "string",
          "readOnly": true
        },
        "workdir": {
          "type": "string",
          "readOnly": true
        }
      }
    },
    "mount_rbd": {
      "description": "mount_rbd describes an RBD mount.  This must have at least and RBD ID associated with it (which becomes the mount's ID), and a provided filesystem type.",
      "type": "object",
      "required": [
        "id",
        "fs_type"
      ],
      "properties": {
        "fs_type": {
          "type": "string"
        },
        "id": {
          "description": "must be a valid rbd device id",
          "type": "integer",
          "format": "int64"
        },
        "mount_options": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "mountpoint": {
          "type": "string",
          "readOnly": true
        },
        "ref": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        }
      }
    },
    "rbd": {
      "description": "rbd describes an RBD map.  To successfully map, at least one monitor, pool and image must be specified.\nAdditionally, you will need options.name and options.secret specified.\n",
      "type": "object",
      "required": [
        "monitors",
        "pool",
        "image"
      ],
      "properties": {
        "id": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        },
        "image": {
          "type": "string",
          "minLength": 1
        },
        "monitors": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "ipv4",
            "pattern": "^(?:[0-9]{1,3}\\.){3}[0-9]{1,3}$"
          }
        },
        "options": {
          "$ref": "#/definitions/rbd_options"
        },
        "pool": {
          "type": "string",
          "minLength": 1
        },
        "refs": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        },
        "snapshot": {
          "type": "string"
        }
      }
    },
    "rbd_options": {
      "type": "object",
      "properties": {
        "abort_on_full": {
          "type": "boolean"
        },
        "alloc_size": {
          "type": "integer"
        },
        "ceph_requires_signatures": {
          "type": "boolean"
        },
        "cephx_sign_messages": {
          "type": "boolean"
        },
        "crc": {
          "type": "boolean"
        },
        "exclusive": {
          "type": "boolean"
        },
        "force": {
          "type": "boolean"
        },
        "fsid": {
          "type": "string"
        },
        "ip": {
          "type": "string",
          "format": "ipv4",
          "pattern": "^(?:[0-9]{1,3}\\.){3}[0-9]{1,3}$"
        },
        "lock_on_read": {
          "type": "boolean"
        },
        "lock_timeout": {
          "type": "integer",
          "format": "int64"
        },
        "mount_timeout": {
          "type": "integer"
        },
        "name": {
          "type": "string"
        },
        "namespace": {
          "type": "string"
        },
        "noceph_requires_signatures": {
          "type": "boolean"
        },
        "nocephx_sign_messages": {
          "type": "boolean"
        },
        "nocrc": {
          "type": "boolean"
        },
        "noshare": {
          "type": "boolean"
        },
        "notcp_nodelay": {
          "type": "boolean"
        },
        "notrim": {
          "type": "boolean"
        },
        "osd_idle_ttl": {
          "type": "integer"
        },
        "osdkeepalive": {
          "type": "integer"
        },
        "queue_depth": {
          "type": "integer"
        },
        "ro": {
          "type": "boolean"
        },
        "rw": {
          "type": "boolean"
        },
        "secret": {
          "type": "string"
        },
        "share": {
          "type": "boolean"
        },
        "tcp_nodelay": {
          "type": "boolean"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http",
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Mange system image containers",
    "title": "Image API",
    "version": "1.0.0"
  },
  "basePath": "/imageapi/v1",
  "paths": {
    "/attach/rbd": {
      "get": {
        "tags": [
          "attach"
        ],
        "operationId": "list_rbds",
        "responses": {
          "200": {
            "description": "list of rbd maps",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/rbd"
              }
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "post": {
        "tags": [
          "attach"
        ],
        "operationId": "map_rbd",
        "parameters": [
          {
            "name": "rbd",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/rbd"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "RBD attach succeed",
            "schema": {
              "$ref": "#/definitions/rbd"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/attach/rbd/{id}": {
      "get": {
        "tags": [
          "attach"
        ],
        "operationId": "get_rbd",
        "responses": {
          "200": {
            "description": "RBD entry",
            "schema": {
              "$ref": "#/definitions/rbd"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "delete": {
        "tags": [
          "attach"
        ],
        "operationId": "unmap_rbd",
        "responses": {
          "204": {
            "description": "Unmapped"
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "integer",
          "format": "int64",
          "name": "id",
          "in": "path",
          "required": true
        }
      ]
    },
    "/container": {
      "get": {
        "description": "Get a list of containers",
        "tags": [
          "containers"
        ],
        "operationId": "list_containers",
        "responses": {
          "200": {
            "description": "List of containers",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/container"
              }
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "post": {
        "description": "Create a container",
        "tags": [
          "containers"
        ],
        "operationId": "create_container",
        "parameters": [
          {
            "name": "container",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/container"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Container creation succeed",
            "schema": {
              "$ref": "#/definitions/container"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/container/{id}": {
      "get": {
        "description": "Get a container definition",
        "tags": [
          "containers"
        ],
        "operationId": "get_container",
        "responses": {
          "200": {
            "description": "Container entry",
            "schema": {
              "$ref": "#/definitions/container"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "delete": {
        "description": "Delete a container defition.  \nThis will stop running containers.\n",
        "tags": [
          "containers"
        ],
        "operationId": "delete_container",
        "responses": {
          "204": {
            "description": "Container deleted"
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "integer",
          "format": "int64",
          "name": "id",
          "in": "path",
          "required": true
        }
      ]
    },
    "/container/{id}/{state}": {
      "get": {
        "description": "Request a (valid) state for a container. \nValid states to request include: ` + "`" + `running` + "`" + `, ` + "`" + `exited` + "`" + `, ` + "`" + `paused` + "`" + ` (paused is not yet implemented)\n",
        "tags": [
          "containers"
        ],
        "operationId": "set_container_state",
        "responses": {
          "200": {
            "description": "Container state changed",
            "schema": {
              "$ref": "#/definitions/container"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "integer",
          "format": "int64",
          "name": "id",
          "in": "path",
          "required": true
        },
        {
          "enum": [
            "running",
            "exited"
          ],
          "type": "string",
          "name": "state",
          "in": "path",
          "required": true
        }
      ]
    },
    "/mount/overlay": {
      "get": {
        "tags": [
          "mounts"
        ],
        "operationId": "list_mounts_overlay",
        "responses": {
          "200": {
            "description": "list of overlay mounts",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/mount_overlay"
              }
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "post": {
        "tags": [
          "mounts"
        ],
        "operationId": "mount_overlay",
        "parameters": [
          {
            "name": "mount",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/mount_overlay"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Overlay mount succeed",
            "schema": {
              "$ref": "#/definitions/mount_overlay"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/mount/overlay/{id}": {
      "get": {
        "tags": [
          "mounts"
        ],
        "operationId": "get_mount_overlay",
        "responses": {
          "200": {
            "description": "Overlay mount entry",
            "schema": {
              "$ref": "#/definitions/mount_overlay"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "delete": {
        "tags": [
          "mounts"
        ],
        "operationId": "unmount_overlay",
        "responses": {
          "204": {
            "description": "Unmounted"
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "integer",
          "format": "int64",
          "name": "id",
          "in": "path",
          "required": true
        }
      ]
    },
    "/mount/rbd": {
      "get": {
        "tags": [
          "mounts"
        ],
        "operationId": "list_mounts_rbd",
        "responses": {
          "200": {
            "description": "list of rbd mounts",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/mount_rbd"
              }
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "post": {
        "tags": [
          "mounts"
        ],
        "operationId": "mount_rbd",
        "parameters": [
          {
            "name": "mount",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/mount_rbd"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "RBD mount succeed",
            "schema": {
              "$ref": "#/definitions/mount_rbd"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/mount/rbd/{id}": {
      "get": {
        "tags": [
          "mounts"
        ],
        "operationId": "get_mount_rbd",
        "responses": {
          "200": {
            "description": "RBD mount entry",
            "schema": {
              "$ref": "#/definitions/mount_rbd"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "delete": {
        "tags": [
          "mounts"
        ],
        "operationId": "unmount_rbd",
        "responses": {
          "204": {
            "description": "Unmounted"
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "integer",
          "format": "int64",
          "name": "id",
          "in": "path",
          "required": true
        }
      ]
    }
  },
  "definitions": {
    "container": {
      "description": "The ` + "`" + `container` + "`" + ` option describes a minimally namespaced container.",
      "type": "object",
      "required": [
        "mount",
        "command"
      ],
      "properties": {
        "command": {
          "type": "string"
        },
        "id": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        },
        "logfile": {
          "type": "string",
          "readOnly": true
        },
        "mount": {
          "$ref": "#/definitions/mount"
        },
        "namespaces": {
          "description": "A list of Linux namespaces to use.\n\nNote: This is currently unused.  All containers currently get ` + "`" + `mnt` + "`" + ` and ` + "`" + `pid` + "`" + `.\n      It's here as a placeholder for future use.\n",
          "type": "array",
          "items": {
            "$ref": "#/definitions/container_namespace"
          }
        },
        "state": {
          "description": "When read, this contains the current container state. On creation, this requests the initial state (valid options: ` + "`" + `created` + "`" + ` or ` + "`" + `running` + "`" + `). The default is ` + "`" + `created` + "`" + `.\n",
          "$ref": "#/definitions/container_state"
        },
        "systemd": {
          "description": "When ` + "`" + `systemd` + "`" + ` is set to ` + "`" + `true` + "`" + `, we will assume that this container will run ` + "`" + `systemd` + "`" + `, and perform the necessary magic dance to make systemd run inside of the container. The default is ` + "`" + `false` + "`" + `.\n",
          "type": "boolean"
        }
      }
    },
    "container_namespace": {
      "description": "Linux namespace",
      "type": "string",
      "enum": [
        "cgroup",
        "ipc",
        "net",
        "mnt",
        "pid",
        "time",
        "user",
        "uts"
      ]
    },
    "container_state": {
      "description": "Valid container states",
      "type": "string",
      "enum": [
        "created",
        "running",
        "stopping",
        "exited",
        "dead"
      ]
    },
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "mount": {
      "description": "Generically address mounts by kind and ID",
      "type": "object",
      "required": [
        "kind",
        "id"
      ],
      "properties": {
        "id": {
          "type": "integer",
          "format": "int64"
        },
        "kind": {
          "type": "string",
          "enum": [
            "overlay",
            "rbd"
          ]
        }
      }
    },
    "mount_overlay": {
      "description": "` + "`" + `mount_overlay` + "`" + ` describes an Overlayfs mount.  All mount points must be RBD ID's.\nAt very least, ` + "`" + `lower` + "`" + ` must be specified.  If ` + "`" + `upper` + "`" + ` length is zero, no ` + "`" + `upper` + "`" + `\nmounts will be used.  ` + "`" + `workdir` + "`" + ` will be assigned automatically.\n\nOverlay mounts are identified by their ` + "`" + `lower` + "`" + ` ID.\n",
      "type": "object",
      "required": [
        "lower"
      ],
      "properties": {
        "id": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        },
        "lower": {
          "description": "This is an array of RBD IDs, interpreted in order for multiple lower dirs. At least one must be specified.",
          "type": "array",
          "items": {
            "type": "integer",
            "format": "int64"
          }
        },
        "mountpoint": {
          "type": "string",
          "readOnly": true
        },
        "ref": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        },
        "upperdir": {
          "description": "currently, upperdir is always a directory in mountDir",
          "type": "string",
          "readOnly": true
        },
        "workdir": {
          "type": "string",
          "readOnly": true
        }
      }
    },
    "mount_rbd": {
      "description": "mount_rbd describes an RBD mount.  This must have at least and RBD ID associated with it (which becomes the mount's ID), and a provided filesystem type.",
      "type": "object",
      "required": [
        "id",
        "fs_type"
      ],
      "properties": {
        "fs_type": {
          "type": "string"
        },
        "id": {
          "description": "must be a valid rbd device id",
          "type": "integer",
          "format": "int64"
        },
        "mount_options": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "mountpoint": {
          "type": "string",
          "readOnly": true
        },
        "ref": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        }
      }
    },
    "rbd": {
      "description": "rbd describes an RBD map.  To successfully map, at least one monitor, pool and image must be specified.\nAdditionally, you will need options.name and options.secret specified.\n",
      "type": "object",
      "required": [
        "monitors",
        "pool",
        "image"
      ],
      "properties": {
        "id": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        },
        "image": {
          "type": "string",
          "minLength": 1
        },
        "monitors": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "ipv4",
            "pattern": "^(?:[0-9]{1,3}\\.){3}[0-9]{1,3}$"
          }
        },
        "options": {
          "$ref": "#/definitions/rbd_options"
        },
        "pool": {
          "type": "string",
          "minLength": 1
        },
        "refs": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        },
        "snapshot": {
          "type": "string"
        }
      }
    },
    "rbd_options": {
      "type": "object",
      "properties": {
        "abort_on_full": {
          "type": "boolean"
        },
        "alloc_size": {
          "type": "integer"
        },
        "ceph_requires_signatures": {
          "type": "boolean"
        },
        "cephx_sign_messages": {
          "type": "boolean"
        },
        "crc": {
          "type": "boolean"
        },
        "exclusive": {
          "type": "boolean"
        },
        "force": {
          "type": "boolean"
        },
        "fsid": {
          "type": "string"
        },
        "ip": {
          "type": "string",
          "format": "ipv4",
          "pattern": "^(?:[0-9]{1,3}\\.){3}[0-9]{1,3}$"
        },
        "lock_on_read": {
          "type": "boolean"
        },
        "lock_timeout": {
          "type": "integer",
          "format": "int64"
        },
        "mount_timeout": {
          "type": "integer"
        },
        "name": {
          "type": "string"
        },
        "namespace": {
          "type": "string"
        },
        "noceph_requires_signatures": {
          "type": "boolean"
        },
        "nocephx_sign_messages": {
          "type": "boolean"
        },
        "nocrc": {
          "type": "boolean"
        },
        "noshare": {
          "type": "boolean"
        },
        "notcp_nodelay": {
          "type": "boolean"
        },
        "notrim": {
          "type": "boolean"
        },
        "osd_idle_ttl": {
          "type": "integer"
        },
        "osdkeepalive": {
          "type": "integer"
        },
        "queue_depth": {
          "type": "integer"
        },
        "ro": {
          "type": "boolean"
        },
        "rw": {
          "type": "boolean"
        },
        "secret": {
          "type": "string"
        },
        "share": {
          "type": "boolean"
        },
        "tcp_nodelay": {
          "type": "boolean"
        }
      }
    }
  }
}`))
}