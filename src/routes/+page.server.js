export async function load({ fetch }) {
    const response = await fetch("http://localhost:8000/fetch-main-page-items");
    const { data } = await response.json();

    return {
        "items": data,
    };
}