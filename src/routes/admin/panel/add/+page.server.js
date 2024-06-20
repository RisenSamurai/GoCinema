export const actions = {
    addActor: async ({ request }) => {
        const formData = await request.formData();

        const name = formData.get("name");
        const lastName = formData.get("lastName");
        const birthday = formData.get("birthday");
        const gender = formData.get("gender");
        const pob = formData.get("pob");
        const biog = formData.get("biog");

        const images = formData.getAll("images");




        const serverFormData = new FormData();

        images.forEach(image => {
            serverFormData.append("images", image);
        })

        serverFormData.append(name, "name");
        serverFormData.append(lastName, "lastName");
        serverFormData.append(birthday, "birthday");
        serverFormData.append(gender, "gender");
        serverFormData.append(pob, "pob");
        serverFormData.append(biog, "biog");




    },
}