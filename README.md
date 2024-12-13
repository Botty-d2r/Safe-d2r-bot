Our Discord: https://discord.gg/qRdArnNQ

<p align="center"> <img src="assets/logo.webp" alt="Discord Community Logo" width="150"> </p> <h3 align="center">Diablo II: Resurrected Bot Community</h3>
This project is a small bot for Diablo II: Resurrected, built for informational and educational purposes only, and not intended for online usage. Feel free to contribute by opening pull requests with new features or bug fixes.

The bot reads game memory and interacts with the game by injecting clicks and keystrokes into the game window. Join our Discord community to report bugs, ask for help, or just to chat: Join Our Discord.

Disclaimer
Can I get banned for using this bot? The answer is yes, you can get banned. We are not responsible for any bans or other consequences that may arise from its use.

Features
Blizzard Sorceress, Nova Sorceress, Hammerdin, and FoH are currently supported
Supported runs: Countess, Andariel, Ancient Tunnels, Summoner, Mephisto, Council, Eldritch, Pindleskin, Nihlathak, Tristram, Lower Kurast, Stony Tomb, The Pit, Arachnid Lair, Baal, Tal Rasha Tombs, Diablo, Cows
Multi-window support (run multiple bots simultaneously)
Bot integration for Discord and Telegram
Companion mode: a leader bot creates games, and other bots join
Pickit based on NIP files
Auto-potion for health and mana (including mercenary)
Chicken when low health
Inventory slot locking
Revive mercenary
CTA buff and class buffs
Auto repair
Skip on immune
Auto-leveling Sorceress and Paladin (WIP)
Auto gambling
Auto cubing (WIP)
Terror Zones (WIP)
Requirements
Diablo II: Resurrected (1280x720 resolution required, windowed mode, with large fonts in accessibility disabled)
Diablo II: LOD 1.13c (IMPORTANT: It will NOT work without it; this step is mandatory)
Quick Start
Preparing the Character
Set up game key bindings for the skills used by your build.
For Blizzard Sorceress, set the left skill to Glacial Spike or Ice Blast, and for Hammerdin, to Blessed Hammer.
Ensure you have TP and ID tomes, one stack of keys, and a key binding for the TP tome.
The Horadric Cube can be kept in the inventory or stashed.
Configure inventory slots for charms.
Running the Tool
Install Diablo II: LOD 1.13c if you haven’t already.
Download the latest release or build it from source.
Extract the ZIP file to a directory of your choice.
Run the executable and follow the setup wizard.
Pickit Rules
Pickit is based on NIP files. These files should be placed in the config/{character}/pickit directory. All .nip files will be loaded, and items are picked, stashed, or sold based on the rules defined in these files.

Development Environment
To build the project from source, you need the following:

Dependencies
Go >= 1.22
Git
Building from Source
shell
Kopier kode
git clone https://github.com/hectorgimenez/project.git
cd project
build.bat
Updating
Fetch the latest changes:

shell
Kopier kode
git pull
build.bat
Note: Customizations in the build directory will be overwritten during updates. Backup files if needed.






