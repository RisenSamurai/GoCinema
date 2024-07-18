export async function load({ fetch }) {
    const response = await fetch("http://localhost:8000/fetch-mainpage-items");
    const items = await response.json();

    return {
        items
    };
}