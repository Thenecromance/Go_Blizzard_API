# Unofficial API 

## Important Notice

1. If you want to use this project to collect data, you need to set up a cookie in player.go -->StringPlayerSummary method. (This is because XXX restricts access to player data through their official API. By using cookies, you can access data just as you log in through a web browser. Please ensure that you abide by the relevant company's terms of service when using this method.)
2. Since the front ends corresponding to different game projects are not in the same group, various methods may vary. Therefore, please use them with caution
3. It is currently in the early stage of development and has not undergone detailed testing


## Prepare to support games 

- World of Warcraft
- World of Warcraft Classic
- Hearthstone
- Warcraft II

## Features

### 📦 Data Models (`/model`)
Comprehensive Go structs with JSON tags for automatic unmarshalling of API responses:
- **Player**: Core character info, attributes, guild, realm, and active spec.
- **Mythic Keystone Profile**: Current mythic rating, best runs, season history, and dungeon details.
- **Equipment**: Detailed breakdown of equipped items, stats, enchantments, set bonuses, and transmog info.
- **Keystone Season**: Seasonal data including affixes and specific run history.

### 🛠 Interfaces (`/Interface`)
Flexible interfaces to allow dependency injection for core infrastructure:
- **Request**: Abstract HTTP client interface (`GET`, `POST`, etc.) to support different networking libraries.
- **Logger**: Abstract logging interface.

### ⚡ Internal Utilities (`/internal`)
- **Token Caching**: In-memory caching mechanism for handling authentication tokens or session identifiers.


## Usage Example

### Parsing Player Data

| Method Start with | Description                      | Return Type          |
|-------------------|----------------------------------|----------------------|
| String**          | Fetch json data (standard)       | string               |
| CN**              | Marshal data (China)             | struct in the same file     |
| BNet**            | Convert Data to BNet Portal data | struct in the model folder  |

## Project Structure

```
/Unofficial_API
├── Interface/       # Core interfaces (Request, Logger)
├── internal/        # Internal logic (Token caching)
├── model/           # Data structures (Player, Equipment, Mythic+, etc.)
├── utils/           # Utility functions (Cache implementation)
├── go.mod           # Module definition
└── README.md        # Documentation
```

## Dependencies

- `github.com/jinzhu/copier`: Used for deep copying structures.

## Disclaimer

This project is an unofficial API wrapper and is not affiliated with, endorsed, sponsored, or specifically approved by Blizzard Entertainment, Inc. World of Warcraft and Blizzard Entertainment are trademarks or registered trademarks of Blizzard Entertainment, Inc. in the U.S. and/or other countries. All data gathered by this tool belongs to their respective owners.

## License

[MIT](LICENSE)