function check(n){
   return n === "tanishk";
}

const p = new Promise((res , rej)=>{
   check("tanishk") ? res("promise resolved") : rej("promise rejected")
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