# [WIP] Go ABM lib

This is a WIP attempt to create a library for Agent-Based Models simulation using pure Go.

Key interfaces:

#### Agent

Code specifying behaviour of the agent:

```go
type Agent interface {
    Run()
}
```
See `models/` for examples.

#### World

Code describing properies of the world:

```go
type World interface {
    Tick() // mark the beginning of the next time period
}
```

See `worlds/` for examples of worlds.

#### UI

Code for representing World with Agents in a visual form:

```go
type UI interface {
    Stop()
    Loop()
}
```

See `ui/` for examples.

## Examples

See `examples/`.

# License

MIT

