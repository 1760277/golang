import http from '../common/http';

class UserService {
    public Login(data: any) {
        return http.post(`/user/login`, data);
    }

    public GetUser() {
        return http.get(`/user`);
    }

    public AddUser(data: any) {
        return http.post(`/user/add`, data);
    }

    public UpdateUser(data: any) {
        return http.post(`/user/update`, data);
    }

    public DeleteUser(id: any) {
        return http.delete(`/user/del/`+ id);
    }
    
}

export default new UserService();
