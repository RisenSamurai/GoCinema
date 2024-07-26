

export async function POST({ request }) {
    const data = await request.FormData();

    const name = data.get('name');
    const year = data.get('year');
    const directors = data.get('directors');
    const releaseDate = data.get('releaseDate');
    const duration = data.get('duration');
    const description = data.get();
    const images = data.getAll("images");
}