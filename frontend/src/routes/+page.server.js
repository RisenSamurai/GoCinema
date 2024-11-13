export async function load({ fetch }) {
    const response = await fetch("http://localhost:8080/tmdb/items");
    const data = await response.json();

    console.log(data)

    return {
        "movies": data.movies,
        "articles": data.articles,
    };
}