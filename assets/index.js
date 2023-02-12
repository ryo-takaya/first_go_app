
function handleOnEvent (e){
    e.preventDefault()

    const body = (document.getElementsByClassName('todo__body')[0]).value
    if (body.length === 0) {
    return
}
    const requestData = "data=" + body
    let request = new XMLHttpRequest();
    request.open('POST', '/add');
    request.setRequestHeader( 'Content-Type', 'application/x-www-form-urlencoded' );
    request.send(requestData)
    request.addEventListener("load", function(){
        const list = document.createElement('li')
        list.innerHTML = this.response
        const ul = document.getElementsByClassName('todo__list')[0]
        ul.appendChild(list);
        document.getElementsByClassName('todo__body')[0].value = ''
}, false);
}
