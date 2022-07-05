// Whole-script strict mode syntax
"use strict";

function displaySelectInput() {
    const search = document.getElementById("search");
    const searchMethod = search.options[search.selectIndex].value;
    console.log(`searchMethod: ${searchMethod}`);
}


function registerListeners() {
    const search = document.getElementById("search");
    search.addEventListener("change", displaySelectInput);
}
registerListeners();
