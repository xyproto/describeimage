# Describe Image

A command line tool that can describe one or more images.

## Requirements

### Run-time requirements

* Ollama (the service must be up and running, and there must be enough memory and CPU and/or GPU available to be able to use the [`llava`](https://ollama.com/library/llava) model).

### Build-time requirements

* Go 1.22 or later

## Installation

    go install github.com/xyproto/describeimage@latest

## Example use

<img align="right" width="150" height="150" alt="Puppy" src="img/puppy.png">

```sh
describeimage img/puppy.png
```

> The image shows a cute puppy sitting on a city sidewalk, looking up at the camera with its tongue out. It's a staged photo that captures the dog in an urban setting, likely intended to evoke feelings of joy and warmth associated with pets in city environments. The project containing this file might be related to pet photography, advertising for pet-friendly services or products, or even a stock image collection focused on animals in urban settings.

## General info

* Version: 1.0.1
* License: BSD-3
* Author: Alexander F. Rødseth &lt;xyproto@archlinux.org&gt;
