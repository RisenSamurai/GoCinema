
export async function load ({ params, fetch}) {

    const { id } = params;

    const response = await fetch(`http://localhost:8080/tmdb/movie/${id}`)

    const data = await response.json();

    return {
        "movie": data,
    };



}