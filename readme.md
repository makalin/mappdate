# mappdate: macOS Updates CLI

A command-line interface tool written in Go that allows you to list and install macOS system updates directly from your terminal. This tool provides an interactive interface to select and manage system updates without leaving the command line.

## Features

- 📋 List all available macOS system updates
- 🔍 View detailed information about each update (version, size, description)
- ✨ Interactive selection menu for choosing updates
- 📦 Install multiple updates in one go
- 🚦 Real-time progress tracking for installations
- ⚠️ Error handling and status reporting

## Prerequisites

- macOS operating system
- Go 1.16 or later
- Administrative privileges (sudo access)

## Installation

1. Clone the repository:
```bash
git clone https://github.com/yourusername/mappdate.git
cd mmappdate
```

2. Install the required dependency:
```bash
go get github.com/manifoldco/promptui
```

3. Build the application:
```bash
go build mappdate.go
```

## Usage

Run the program with sudo privileges:

```bash
sudo ./updates
```

The program will:
1. Check for available updates
2. Display an interactive menu with all available updates
3. Allow you to select multiple updates using space bar
4. Install selected updates with progress tracking

### Interactive Menu Controls

- `↑`/`↓`: Navigate through updates
- `space`: Select/deselect an update
- `enter`: Confirm selection and start installation
- `ctrl+c`: Exit the program

## Example Output

```
Available Updates:
* macOS Ventura 13.4.1 (Version: 13.4.1, Size: 3.2GB)
  A critical update for your Mac that includes security fixes...

* Safari 16.5.2 (Version: 16.5.2, Size: 89.3MB)
  Important security and stability updates...

Select updates to install (use space to select, enter to confirm)
  > [x] macOS Ventura 13.4.1
    [ ] Safari 16.5.2

Installing selected updates...
Installing macOS Ventura 13.4.1...
Successfully installed macOS Ventura 13.4.1
```

## Error Handling

The program includes comprehensive error handling for common scenarios:
- Insufficient privileges
- Network connectivity issues
- Invalid update selections
- Installation failures

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Uses the `softwareupdate` command-line tool provided by Apple
- Interactive menu powered by [promptui](https://github.com/manifoldco/promptui)

## Security Note

This application requires sudo privileges to function properly as system updates on macOS can only be installed with administrative access. Always review the source code of applications that require sudo privileges before running them.

## Support

If you encounter any issues or have questions, please file an issue in the GitHub repository.