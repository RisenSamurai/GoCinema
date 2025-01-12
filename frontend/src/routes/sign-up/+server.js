
export async function POST({ request }) {
    const body = await request.json();

    const {email, password} = body;


   console.log('Password:', password);
   console.log('Email:', email);

   let SERVER_ADDRESS = 'localhost:8080';

   const response = await fetch(`http://${SERVER_ADDRESS}/auth/sign-up`, {
       method: 'POST',
       headers: {
           'Content-Type': 'application/json',
       },
       body: JSON.stringify({email, password}),
   })
    const data = await response.json();

    console.log(data);


   return new Response(JSON.stringify({success: true, message: `Data was sent successfully`}, {
       status: 200,
       headers: { 'content-type': 'application/json' },
   }));



}