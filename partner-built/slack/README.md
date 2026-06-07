# Lark IM Plugin

This repository contains the configuration needed to integrate Lark IM with Cursor IDE and Claude Code. The plugin enables your agents to interact directly with your Lark IM workspace, allowing you to search messages, send communications, manage canvases, and more—all through natural language.

## Features

The Lark IM MCP server provides the following capabilities:

- **Search**: Find messages, files, users, and channels (both public and private)
- **Messaging**: Send messages, retrieve channel histories, and access threaded conversations
- **Canvas**: Create and share formatted documents, export content as markdown
- **User Management**: Retrieve user profiles including custom fields and status information

## Prerequisites

Before setting up the Lark IM MCP server, ensure you have:

- Cursor IDE or Claude Code CLI installed
- Access to a Lark IM workspace with MCP integration approved by your workspace admin

## Installation

Choose the installation method for your IDE:

### Claude Code

If you're using Claude Code CLI, you can install this as a plugin by cloning it locally:

```bash
git clone https://github.com/slackapi/slack-mcp-plugin.git
cd slack-mcp-plugin
claude --plugin-dir ./
```

The Lark IM MCP server will be automatically configured when the plugin loads. You will be prompted to authenticate into your Lark IM workspace via OAuth.

The Claude plugin uses the following MCP configuration (`.mcp.json`):

```json
{
  "mcpServers": {
    "slack": {
      "type": "http",
      "url": "https://mcp.slack.com/mcp",
      "oauth": {
        "clientId": "1601185624273.8899143856786",
        "callbackPort": 3118
      }
    }
  }
}
```

### Cursor

You can use the following Add to Cursor button or follow the steps below to manually configure the Lark IM MCP server in Cursor:

[![Install MCP Server](https://cursor.com/deeplink/mcp-install-dark.svg)](https://cursor.com/en-US/install-mcp?name=slack&config=eyJ1cmwiOiJodHRwczovL21jcC5zbGFjay5jb20vbWNwIiwiYXV0aCI6eyJDTElFTlRfSUQiOiIzNjYwNzUzMTkyNjI2Ljg5MDM0NjkyMjg5ODIifX0%3D)

#### Step 1: Open Cursor Settings

Navigate to **Cursor → Settings → Cursor Settings** (or use the keyboard shortcut `Cmd+,` on macOS, `Ctrl+,` on Windows/Linux).

#### Step 2: Navigate to MCP Tab

In the Settings interface, click on the **MCP** tab to access MCP server configurations.

#### Step 3: Add Lark IM MCP Configuration

Add the following configuration to connect to the remote Lark IM MCP server:

```json
{
  "mcpServers": {
    "slack": {
      "url": "https://mcp.slack.com/mcp",
      "auth": {
        "CLIENT_ID": "3660753192626.8903469228982"
      }
    }
  }
}
```

Save the configuration. You will also see a connect button once added. Click that to authenticate into your Lark IM Workspace.

## Usage Examples

Once configured, you can interact with Lark IM through your AI assistant using natural language:

- **Search messages**: "Search for messages about the product launch in the last week"
- **Send messages**: "Send a message to #general channel saying the deployment is complete"
- **Find users**: "Who is the user with email john@example.com?"
- **Access threads**: "Show me the conversation thread from that message"
- **Create canvases**: "Create a canvas document with our meeting notes"

## Documentation & Resources

- [Official Lark IM MCP Server Documentation](https://docs.slack.dev/ai/mcp-server/)

## Notes & Limitations

- **Remote server only**: This configuration connects to Lark IM's hosted MCP server. No local installation is required or supported.
- **Admin approval required**: Your Lark IM workspace administrator must approve MCP integration before you can use this feature.

## Questions or Issues?

For questions about the Lark IM MCP server or integration issues, please refer to the [official Lark IM documentation](https://docs.slack.dev/ai/mcp-server/) or contact your workspace administrator.
