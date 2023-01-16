# SMCNTR

Here are the source files for my first Youtube Short. It is a contrived example on container minimalism based on [multi-stage builds](https://docs.docker.com/build/building/multi-stage/). It is contrived in that it uses the `scratch` container, which holds no external dependencies and even requires go to build the static binary with its own network stack. In a more realistic setting you'd probably want to use a lightweight base container like busybox, especially if you plan on including CA tooling, etc. But this drives the point home: you are probably including a bunch of stuff you don't need (and shouldn't include), which puts a burden on your network and potentially increases your attack surface. 

You can watch the video here: https://www.youtube.com/shorts/BgWhdBGn5qA
