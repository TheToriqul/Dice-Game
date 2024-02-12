function roll() {
    const minValue = 1;
    const maxValue = 6;
    return Math.floor(Math.random() * (maxValue - minValue + 1)) + minValue;
  }
  
  let players = 0;
  while (players < 2 || players > 4) {
    players = Number(prompt("Enter the number of players (2 - 4): "));
    if (isNaN(players)) {
      console.log("Invalid input. Please enter a number.");
    } else if (players < 2 || players > 4) {
      console.log("Must be between 2 - 4 players.");
    }
  }
  
  const maxScore = 50;
  const playerScores = new Array(players).fill(0);
  
  while (Math.max(...playerScores) < maxScore) {
    for (let playerIdx = 0; playerIdx < players; playerIdx++) {
      console.log("\nPlayer number", playerIdx + 1, "turn has just started!");
      console.log("Your total score is:", playerScores[playerIdx], "\n");
      let currentScore = 0;
  
      while (true) {
        const shouldRoll = prompt("Would you like to roll (y)? ");
        if (shouldRoll.toLowerCase() !== "y") {
          break;
        }
  
        const value = roll();
        if (value === 1) {
          console.log("You rolled a 1! Turn done");
          currentScore = 0;
          break;
        } else {
          currentScore += value;
          console.log("You rolled a: ", value);
        }
  
        console.log("Your score is:", currentScore);
      }
  
      playerScores[playerIdx] += currentScore;
      console.log("Your total score is:", playerScores[playerIdx]);
    }
  }
  
  const winningIdx = playerScores.indexOf(Math.max(...playerScores));
  console.log("Player number", winningIdx + 1, "is the winner with a score of:", Math.max(...playerScores));
  