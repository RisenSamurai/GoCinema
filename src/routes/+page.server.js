export async function load({ fetch }) {
    const response = await fetch("http://localhost:8000/get-movies");

}