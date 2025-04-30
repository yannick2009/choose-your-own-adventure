## Description

The "Choose Your Own Adventure" game is based on the format of the popular book series that allows readers to make choices that determine the main character's actions and the story's outcome. In this implementation:

- The story is read from a JSON file
- The application serves web pages with different story chapters
- Users make choices by clicking on links
- Each choice leads to a different part of the story

## Features

- Dynamic HTML generation based on story data
- Custom story file support
- Responsive web design
- Configurable template handling

## Installation

```bash
git clone https://github.com/yannick2009/choose-your-own-adventure.git
cd choose-your-own-adventure
go mod init choose-your-own-adventure
go build
```

## Usage

1. Start the server:

```bash
./choose-your-own-adventure
```

2. Open your browser and navigate to `http://localhost:3000`
3. Follow the story and make choices to explore different paths

## Customization

You can provide your own story file in JSON format:

```bash
./choose-your-own-adventure -path=mystory.json
```

## Learning Objectives

- Working with JSON in Go
- Creating HTTP handlers
- Using Go templates
- Designing web applications with Go

## License

This project is licensed under the MIT License - see the LICENSE file for details.
