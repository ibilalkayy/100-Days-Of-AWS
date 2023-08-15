const newItem = document.getElementById("new-item");
const addBtn = document.getElementById("add-btn");
const todoList = document.getElementById("todo-list");

addBtn.addEventListener("click", addItem);
todoList.addEventListener("click", deleteItem);

function addItem(event) {
  event.preventDefault();
  if (newItem.value !== "") {
    const li = document.createElement("li");
    const itemText = document.createTextNode(newItem.value);
    li.appendChild(itemText);
    todoList.appendChild(li);
    newItem.value = "";
  }
}

function deleteItem(event) {
  const target = event.target;
  if (target.classList.contains("delete-btn")) {
    const li = target.parentNode;
    todoList.removeChild(li);
  } else if (target.tagName === "LI") {
    target.classList.toggle("checked");
  }
}
