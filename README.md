## Name of the game: Take the Crown

**Description:** Take the Crown is a dice rolling game for 2-4 players. Players take turns rolling a die, and their score for each turn is the sum of the values they roll. If a player rolls a 1, their turn ends and their score for that turn is 0. The first player to reach a score of 50 or more wins the game.

Here are the steps on how to play the game:

1. Determine how many players will be playing (2-4).
2. Each player starts with a score of 0.
3. The first player rolls the die.
4. If the player rolls a 1, their turn ends and their score for that turn is 0. Otherwise, the value they rolled is added to their score for the turn.
5. The player can choose to roll again or end their turn.
6. If the player chooses to roll again, steps 4 and 5 are repeated.
7. If the player chooses to end their turn, their score for the turn is added to their total score.
8. Play continues to the next player in clockwise order.
9. The first player to reach a total score of 50 or more wins the game.

I hope you enjoy playing Take the Crown!


## Take the Crown Code Breakdown for Beginners like me:

This code simulates a turn-based dice-rolling game called "Take the Crown" for 2-4 players. Let's break it down step-by-step:

**1. Setting Up:**

* `roll` function: Simulates rolling a die by generating a random number between 1 and 6.
* `while True` loop: Ensures user enters a valid number of players (2-4).
* `max_score` variable: Sets the winning score to 50.
* `player_score` list: Tracks each player's current score with initial 0s.

**2. Gameplay Loop:**

* `while max(player_score) < max_score` loop: Continues until someone wins.
* `for player_idx in range(players)` loop: Cycles through each player's turn.
* `current_score` variable: Tracks the score for the current player's turn.

**3. Player Turn:**

* `print` statements inform the player about their turn and score.
* `while True` loop: Allows the player to keep rolling until they choose to stop.
* `should_roll` input asks if the player wants to roll again ("y" for yes).
* `if value == 1`: If they roll a 1, their turn ends with a score of 0.
* `else`: Add the rolled value to `current_score` and inform the player.
* `player_score` list updates the player's total score with `current_score`.

**4. Finding the Winner:**

* `max_score` updates to the highest score among players.
* `winning_idx` finds the index of the player with the highest score.
* `print` statements announce the winner and their score.

**Remember:**

* This code simulates the game logic, not the graphical part.
* You can customize the code to change things like the winning score, player names, or add new rules.

**Have fun playing Take the Crown!**