let arr = ["a","a","b","c","d"]

let count = [...new Set(arr)]



console.log(arr.filter(x => x != count[x]))