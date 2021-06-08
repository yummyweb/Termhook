# Termhook

Termhooks are user defined callbacks in the terminal. They are similar to webhooks but instead of a url, declarative configuration is used. Termhook can integrate with your error logging tool, a mail delivery service or your own product.

## Termhook Callbacks

For now the callbacks include:

- Start: Occurs when the process first starts.

- Interrupt: Occurs when CTRL+C is pressed, or SIGINT signal is called.

- End: Occurs after the process ends.

- TError: Occurs when Termhook runs into an error.

## Termhook API

Using the Termhook API, you can seamlessly integrate your product with Termhook and get started. The Termhook API allows you to create functions, callable by your users.
