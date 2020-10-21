
class AuthService {
    public isAuthenticated() {
        if (this.getToken()) {
            return true;
        } else {
            return false;
        }
    }

    private getToken() {
        return localStorage.getItem('demotoken');
    }
}

export default new AuthService();
