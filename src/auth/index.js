import auth0 from "auth0-js";

const webAuth = new auth0.webAuth({
	domain: "dev-s9lusa35.us.auth0.com",
	clientId: "aJZSL2BB5IWu4EEDjeJYM80azlRQPicA",
	redirectUri: '',
	responseType: 'token id_token',
	scope: "openid profile"
});

const login = () =>{
	webAuth.authorize();
}

export { login };