<div align="center">

# Simple Event Stream

<p>
    <img src="https://img.shields.io/static/v1?label=NATS&message=Latest&style=flat&logo=nats&colorA=24273A&colorB=8AADF4&logoColor=CAD3F5"/>
    <img src="https://img.shields.io/static/v1?label=Go&message=1.20&style=flat&logo=go&colorA=24273A&colorB=00ADD8&logoColor=CAD3F5"/>
    <img src="https://img.shields.io/static/v1?label=Node.js&message=22.6&style=flat&logo=node.js&colorA=24273A&colorB=339933&logoColor=CAD3F5"/>
</p>

A simple event-driven system using [NATS](https://nats.io/) for real-time messaging between a publisher and a subscriber. This project demonstrates basic message publishing and subscribing with a focus on simplicity.

</div>

## Overview

This project consists of these main components:

1. **NATS Server**: A lightweight and high-performance messaging system.
2. **Publisher**: A Go-based service that publishes UUID messages to a NATS subject.
3. **Subscriber**: A Node.js-based service that subscribes to messages from the NATS subject.

The publisher generates UUID messages every second and sends them to a NATS subject called `updates`. The subscriber listens to the `updates` subject and logs any messages received.

##

<p align="center"><img src="https://imgur.com/nbkZw8B.png" width=1000px></p>
