# DescribeImage

Use LLMs that are running locally (or on a server defined in `OLLAMA_HOST`) to describe the given images.

This repository contains a command line utility that can be used to describe images with Ollama.

## Requirements

### Run-time requirements

* Ollama (the service must be up and running, and there must be enough memory and CPU and/or GPU available to be able to use the user-configured LLM model for vision tasks (for example the [`llava`](https://ollama.com/library/llava) model).
* [`llm-manager`](https://github.com/xyproto/llm-manager) can be used to configure which model to use for the `vision` task.

### Build-time requirements

* Go 1.22 or later

## Installation

    go install github.com/xyproto/describeimage@latest

The executable ends up in `~/go/bin` unless Go has been configured to place it somewhere else.

## Example use

```sh
describeimage puppy.png
```

<img align="right" width="150" height="150" alt="Puppy" src="img/puppy.png">

> The image shows a cute puppy sitting on the sidewalk in what appears to be an urban setting. The puppy is looking directly at the camera with its tongue out, giving a playful and happy expression. It has a fluffy coat of light brown fur. In the background, there's a cityscape with tall buildings, street lamps, and a clear sky. The lighting suggests it might be late afternoon or early evening, as indicated by the warm glow on the buildings. There are no visible texts in the image.

```sh
describeimage meloncat.jpg
```

<img align="right" width="150" height="150" alt="Melon Cat" src="img/meloncat.jpg">

> The image shows a cat with a humorous and unusual appearance. The cat's head is covered by a large, green watermelon slice that has been cut to fit snugly over its head. The cat's eyes are visible through the watermelon, and it appears to be looking directly at the camera with a somewhat disgruntled or unimpressed expression. The background is simple and does not distract from the cat and its watermelon headpiece.

(note that it's really a lime)

## General info

* Version: 1.3.2
* License: BSD-3
* Author: Alexander F. Rødseth &lt;xyproto@archlinux.org&gt;
