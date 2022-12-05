# Capstone Proposal
## Summary

This capstone project's goal is to create a bot that
can connect to a discord server, read the messages posted
in that server and use those messages as well as messages
sent by the bot to create a game.

This project will involve connecting to the Discord API
using webhooks and the discordgo package, as well as the
Discord Developer Portal to create a private token that
can be used for the bot to access the code. I do not want this
token to be publically accessible so I will be using a .env file to contain this token
and a .gitignore to ensure it is not accessible on github.

For this project I will demonstrate the following:

1. Building a REST client API in Go that can interact with the Discord API
2. Creating a handler function in GO that can send requests to the Discord API when certain criteria are met in order to create a game
3. Pushing files to github and following PR workflows 
4. Using .env files to obscure sensitive information
5. Using .gitignore to obscure sensitive information
6. Using structs in Go as objects within a game
7. Following good coding practices (i.e. handling all errors, and logging them accordingly)


## User Stories

### As a user, I would like to invite this Discord bot to my Discord Server and have it initiate a game with another user in my server.

**Acceptance Criteria**

Given the bot is invited to your Discord Server, the bot is running, and
a user types in a prompt in the correct format, the bot will initiate a game
between two users

Example:
```
User1: fight @User2
```

Once a user types in this prompt into a Discord Channel that the bot has permission to read, the bot
would then initiate a game between the author of the prompt and the user mentioned in the message.

The bot would then respond with a confirmation that the initiating user is requesting to fight the recipient.

Example: 
```
FightBot: @User1 is requesting to fight @User2
```

### As a user, I would like to then play the game I initiated with my friends

**Acceptance Criteria**

The bot will flip a coin to see who goes first (or since we are using code, it will generate a random integer, either a 0 or a 1). Then, the bot will prompt the user to let them know it is their turn.

Example:
```
Fightbot: @User1 it's your turn.
```

A user will then type in "punch" and the bot will treat this as an attack on the other user and respond accordingly.
The amount of damage that the each attack does will be a random integer between 0 and 60.

Example:
```
Fightbot: you just punched @User2 for 10 Damage, leaving them with 90 HP
```
The bot would then switch turns and prompt the second user to punch and this would repeat until one of the users 
is left with 0 HP, at which point the bot would declare a winner and end the game.

Example: 
```
Fightbot: you just punched @User1 for 10 Damage, leaving them with 0 HP.

Fightbot: @User2 is your winner of the round!
```
