# DescribeImage

Use LLMs that are running locally (or on a server defined in `OLLAMA_HOST`) to describe the given images.

This repository contains a command line utility that can be used to describe images with Ollama.

## Requirements

### Run-time requirements

* Ollama (the service must be up and running, and there must be enough memory and CPU and/or GPU available to be able to use the user-configured LLM model for vision tasks (for example the [`llava`](https://ollama.com/library/llava) model).
* [`llm-manager`](https://github.com/xyproto/llm-manager) can be used to configure which model to use for the `vision` task (the `-m` flag can also be used).

### Build-time requirements

* Go 1.22 or later

## Installation

    go install github.com/xyproto/describeimage@latest

The executable ends up in `~/go/bin` unless Go has been configured to place it somewhere else.

## Models

The default vision model is `qwen2.5vl:3b` (~3.5 GB).

For machines with >4G VRAM, these might also give good results:

* `gemma3:4b` (~3.3 GB)
* `llava-llama3:8b` (~5.5 GB)
* `llama3.2-vision:latest` (~7.9 GB)

Set the `vision` model in `~/.config/llm-manager/llm.conf` or `/etc/llm.conf`, or pass it with the `-m` flag.

## Example use

```sh
describeimage meloncat.jpg
```

<img align="left" width="170" height="170" alt="Melon Cat" src="img/meloncat.jpg">

> The image shows a cat with a humorous and unusual appearance. The cat's head is covered by a large, green watermelon slice that has been cut to fit snugly over its head. The cat's eyes are visible through the watermelon, and it appears to be looking directly at the camera with a somewhat disgruntled or unimpressed expression. The background is simple and does not distract from the cat and its watermelon headpiece.

(note that it's really an oversized lime)

---

<img align="right" width="400" height="267" alt="Horses" src="img/horses.png">

> The image depicts a serene pastoral scene with a group of horses grazing in a lush green field. In the background, majestic mountains rise against a clear blue sky, partially covered by a few scattered clouds. The landscape is framed by a line of trees, adding depth and contrast to the scene. The overall atmosphere is peaceful and idyllic, capturing the beauty of nature and the tranquility of rural life.

---

## General info

* Version: 1.4.0
* License: BSD-3
* Author: Alexander F. Rødseth &lt;xyproto@archlinux.org&gt;
