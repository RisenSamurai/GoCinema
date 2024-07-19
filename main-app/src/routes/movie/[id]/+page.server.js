
export async function load ({ params, fetch}) {

    const { id } = params;

    const response = await fetch(`http://localhost:8000/fetch-movie/${id}`)


    if (response.ok) {
        const { movie } = await response.json();
        return {
            movie: movie,
        }
    } else {
        throw new Error(`Could not find movie with id ${id}`);
    }


}