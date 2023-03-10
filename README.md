# About

Even though you can use this CLI to generate some cool fun images and show off in front of your friends and colleagues,
but this is mostly an example to show one of the many potentials in using OpenAI and integrating with it!

Using this command, you can easily generate unique cool images in a few seconds and share them on Slack, Twitter, 
or any other social media! 😎

# An example video to show how it works

[![An example command to generate random images using OpenAI](https://img.youtube.com/vi/Ep66K0h0n8Y/0.jpg)](https://youtu.be/Ep66K0h0n8Y)

# Prerequisites

- [Download an install Go ](https://go.dev/doc/install) locally
- Alternatively you can [install and use Docker](https://docs.docker.com/get-docker/) if you are not willing to install
  Go locally

# How to use

- Sign up for OpenAI
- Create an [API key](https://beta.openai.com/account/api-keys)
- Clone the repo
- Create a file named `token` inside the `secrets` directory
```bash
touch ./secrets/token
```
- Paste the value of your API key in the `token` file
- Alternatively, you can set the `OPENAI_API_KEY` environment variable with the value of your API key.

### If Go is installed locally

- Run the command

```bash
# Example:
go run main.go -p "Dancing cat singing in a digital ocean" -s "1024x1024"
```

- Click on the link provided in the output and see your generated image!

### If Go is not installed locally, but Docker is installed

- Run the command

```bash
# Example:
docker run -it --rm -v $(pwd):/app -w /app golang:1.16 go run main.go -p "Dancing cat singing in a digital ocean" -s "1024x1024"
```

# How to make it even easier to use

You can even build the project and create an alias on your laptop to call this command from anywhere in your terminal!

```bash
# 1. Build the project (run it while you are in this path)
go build -o open-ai-image-generate

# 2.1. You can create an alias like this (if you prefer not to build the project)
alias 'open_ai_generate_image=docker run -it --rm -v /path/to/cloned/repo:/app -w /app golang:1.16 go run main.go'

# 2.2. You can create an alias like this (if you have also built the project)
alias 'open_ai_generate_image=/path/to/cloned/repo/open_ai_generate_image'

# 3. Then you can run it like this
open_ai_generate_image -p "Dancing cat singing in a digital ocean" -s "1024x1024"
```

Alternatively, you can simply either copy the executable to your /usr/local/bin or /usr/bin or anywhere else
defined in your `$PATH` and call it from anywhere in your terminal, or you can add the project's directory
to your `$PATH`.

Also, alternatively, you can run `go install` to install the project in your `$GOPATH` and call it from anywhere.
If you do this, remember to:
- Add your `$GOPATH` to your `$PATH` if you haven't already.
- Set the `OPENAI_API_KEY` environment variable with the value of your API key.

# Parameters

- `-p` or `--prompt` is the prompt that tells OpenAI what image to generate
- `-s` or `--size` is the size of the image to generate

# Limitations

Based on the OpenAI's [documentation](https://beta.openai.com/docs/api-reference/images):

- Prompt should be less than **1000** characters.
- Size **must be** one of these values: `256x256`, `512x512`, or `1024x1024`

# Important notes

- Please do not use inappropriate words in the prompt.
- This program is not intended to be used by children.
- Never share any of your API keys with anyone or commit them even to your private repos!

# License
This project is licensed under the MIT License - see the [LICENSE.md](./LICENSE.md) file for details.
