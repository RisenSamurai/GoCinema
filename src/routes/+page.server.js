export async function load({ fetch }) {
    const respone = await fetch("http://localhost:8000/fetch-mainpage-items");
    const items = await respone.json();

    return {
        items
    };
}