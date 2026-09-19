function check(n){
   return n === "tanishk";
}

const p = new Promise((res , rej)=>{
    res("hello world")
    rej("hello world rejected ")
})

fetchdata = async()=>{
    try{
        console.log(await p)
    }catch(e){
        console.log(e)
    }
}
fetchdata()

console.log("hello world")


// promisified readFile: manual wrapper
const fs = require("node:fs");
const { promisify } = require("node:util");

const readFileP = (path) =>
    new Promise((resolve, reject) =>
        fs.readFile(path, "utf8", (err, data) => (err ? reject(err) : resolve(data)))
    );

// same thing via util.promisify
const readFileAsync = promisify(fs.readFile);

readFileP(__filename).then((d) => console.log("manual:", d.length));
readFileAsync(__filename, "utf8").then((d) => console.log("promisify:", d.length));
