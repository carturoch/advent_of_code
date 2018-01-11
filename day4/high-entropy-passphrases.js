function uniqueWords(sen) {
  return sen.split(" ").reduce((acc, w) => {
    // const fn = (uw) => uw == w; // Step 1
    const fn = (uw) => areAnagrams(uw, w);
    if (acc.findIndex(fn) == -1 && w != "") {
      acc.push(w);
    }
    return acc;
  }, []);
}

function areAnagrams(a, b) {
  return a.split("").sort().join("") == b.split("").sort().join("");
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
  areAnagrams: areAnagrams,
  uniqueWords: uniqueWords
};

// const fs = require("fs");
// fs.readFile("./phrases.txt", "utf8", (err, content) => {
//   console.info(content.split("\n").filter(isValidPassphrase).length);
// })