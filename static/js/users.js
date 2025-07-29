/**
 *
 * @param {Number} id
 */
async function deleteUser(id) {
  const res = await fetch(`/admin/users/${id}`, {
    method: "DELETE",
  });
  if (res.status >= 200 && res.status <= 300) {
    document.getElementById(`row-user-${id}`).remove();
  } else {
    console.error(`Response Status: ${res.status}`);
    console.error(`Failed to delete User {id:${id}}`);
  }
}

async function shareUser(id) {
  const res = await fetch(`/admin/users/${id}/share`);
  if (res.status >= 200 && res.status <= 300) {
    const data = await res.text()
    console.log(data);
  } else {
    console.error(`Response Status: ${res.status}`);
    console.error(`Failed to get User share link`);
  }
}
