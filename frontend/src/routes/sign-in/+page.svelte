<script>
    import { isLogged } from "../store.js";
    let message = $state('');

    let email = $state('');
    let password = $state('');
    let confirmPassword = $state('');



   async function handleSubmit(e) {
        e.preventDefault();


        if(password === confirmPassword) {

            try {

                const request = await fetch("/sign-in", {
                    method: "POST",
                    body: JSON.stringify({email, password, confirmPassword}),
                    headers: { "Content-Type": "application/json" },
                });

                const data = await request.json();

                if(request.status === 200) {
                    const JWT = data.token;
                    sessionStorage.setItem("JWT", JWT);
                    isLogged.set(true);

                } else {
                    console.log("Error logged in!", data.error);

                }

            } catch (e) {
                console.error(e.message);
                $message = e.message;
                throw e;
            }


        }
    }
</script>



<div class="flex flex-col bg-cinema-secondary justify-center h-screen text-white self-center">
    {#if message}<span>{message}</span>{/if}
    <form onsubmit={handleSubmit} class="flex flex-col justify-center self-center">

        <label for="email">Email</label>
        <input class="text-black" type="email" id="email" name="email" bind:value={email} required/>
        <label for="password">Password</label>
        <input class="text-black" type="password" id="password" name="password" bind:value={password} required/>
        <label for="password_repeat">Confirm Password</label>
        <input class="text-black" type="password" id="password_repeat" name="password_repeat" bind:value={confirmPassword} required/>

        <button type="submit" class="btn btn-primary">Submit</button>
    </form>

    <span class="flex self-center">OR</span>

    <div class="flex self-center">
        <span>Google</span>
    </div>
</div>