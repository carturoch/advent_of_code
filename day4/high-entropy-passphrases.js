function uniqueWords(sen) {
  return sen.split(" ").reduce((acc, w) => {
    if (acc.indexOf(w) == -1 && w != "") {
      acc.push(w);
    }
    return acc;
  }, []);
}

function isValidPassphrase(sen) {
  if (sen == "") {
    return false;
  }
  let totalWords = sen.split(" ").filter(w => w != "").length;
  let totalUniqWords = uniqueWords(sen).length;
  return totalWords == totalUniqWords;
}

module.exports = {
  isValidPassphrase: isValidPassphrase,
  uniqueWords: uniqueWords
};

// const fs = require("fs");
// fs.readFile("./phrases.txt", "utf8", (err, content) => {
//   console.info(content.split("\n").filter(isValidPassphrase).length);
// })