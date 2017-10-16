const list = document.getElementById('list')
const itemlist = Array.from(list.children)
let selectedItems = []

list.addEventListener('click', e => {
  let target = e.target
  if (target.nodeName === 'INPUT') {
    e.target.checked = !e.target.checked
  }
  while (target.nodeName !== 'TR') {
    target = target.parentNode
  }
  target.cells[0].children[0].checked = !target.cells[0].children[0].checked
  if (target.cells[0].children[0].checked) {
    if(!selectedItems.includes(target.id)) {
      selectedItems.push(target.id)
    }
  } else {
    const index = selectedItems.indexOf(target.id)
    if (index !== -1) {
      Array.prototype.splice.call(selectedItems, index, 1)
    }
  }
  allCheckBox.checked = selectedItems.length === itemlist.length
})
