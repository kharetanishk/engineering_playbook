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