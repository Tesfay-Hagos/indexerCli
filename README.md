# indexerCli

This is a command-line interface (CLI) designed for the kwil Blockchain indexer. The CLI allows users to interact with the kwilindexer by executing various commands.

## Usage

To use this CLI, follow the steps below:

1. Download the repository from [GitHub](https://github.com/Tesfay-Hagos/indexerCli).
2. Store this repository and the kwilindexer in the same folder. Both must be extracted. For example, create a directory with the name "indexer" and store the extracted kwilindexer repository and extracted kwilcli repository in this folder.
3. Go to the working directory of indexercli and execute the following command: `go build -o ../kwilindexer/kwil`. This will build the CLI executable and place it in the kwilindexer folder.
4. You can now use the commands of the CLI to interact with the kwilindexer.

## Commands

The following commands are available:

- `./kwil -l` or `./kwil --list`: Lists all available configurations.
- `./kwil --help` or `./kwil -h`: Displays details on command usage and flags.
- `./kwil --sets`: Sets values for the configuration.

## Flags

The following flags can be used with the `./kwil --sets` command to set specific values for the configuration:

- `-c`: Sets the PG-connection value. Example: `-c "postgres://indexer:indexer123@localhost:5433/indexer?sslmode=disable"`.
- `-e`: Sets the commetbft_endpoint value. Example: `-e "https://localhost:26657"`.
- `-a`: Sets the ListenAddr value. Example: `-a ":8000"` or `-a "127.0.0.1:8000"`.
- `-f`: Sets the poll frequency to fetch node updates. Example: `-f 5` (5 seconds).
- `-b`: Sets the MaxBlockPagination value. Example: `-b 30` (30 seconds).
- `-x`: Sets the MaxTxPagination value. Example: `-x 30` (30 seconds).
