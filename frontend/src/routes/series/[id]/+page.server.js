
export async function load ({ params, fetch}) {

    const { id } = params;

    const response = await fetch(`http://localhost:8080/tmdb/series/${id}`)

    const data = await response.json();

    return {
        "movie": data,
    };



}