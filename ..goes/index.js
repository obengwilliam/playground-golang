async function do1 () {
  var b = await new Promise((resolve, _) => {
    resolve(2)
  })
  return b
}

;(async function a () {
  console.log(await do1())
})()
