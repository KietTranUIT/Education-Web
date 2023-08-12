function submitForm(event) {
    event.preventDefault();
  
    // Lấy dữ liệu từ các trường form
    var formData = {
      username: document.getElementById('username').value,
      password: document.getElementById('password').value,
      surname: document.getElementById('surname').value,
      name: document.getElementById('name').value,
      gender: document.querySelector('input[name="gender"]:checked').value,
      dateofbirth: document.getElementById('date_of_birth').value,
      placeofbirth: document.getElementById('place_of_birth').value
      // Thêm các trường khác tùy ý
    };
  
    // Chuyển đổi object JSON thành chuỗi JSON
    var jsonData = JSON.stringify(formData);
  
    // Gán giá trị cho trường ẩn jsonData
    //document.getElementById('jsonData').value = jsonData;
  
    // Gửi yêu cầu POST
    fetch('/student/signup/sign_up_student', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: jsonData
    })
      .then(response => response.json())
      .then(jsonData => {
        // Xử lý dữ liệu JSON ở đây
        console.log('Received data:', jsonData);
      })
      .catch(error => {
        // Xử lý lỗi
      });
  }

function AccessLoginStudentPage() {
  fetch('/student/login', )
  .then(response => {
    if (response.ok) {
      window.location.href = response.url;
    } else {
      console.error('Yêu cầu không thành công. Mã trạng thái HTTP:', response.status);
    }
  })
  .catch(error => {
    // Xử lý lỗi ở đây
     console.error('Lỗi khi yêu cầu trang mới:', error);
  });
}