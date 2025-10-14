const canvas = document.getElementById("anime-bg");
const ctx = canvas.getContext("2d");

canvas.width = window.innerWidth;
canvas.height = window.innerHeight;

let particlesArray = [];
const mouse = { x: null, y: null };

class Particle {
    constructor(x, y, idle = false) {
        this.x = x || Math.random() * canvas.width;
        this.y = y || Math.random() * canvas.height;
        this.size = Math.random() * 5 + 1;
        this.speedX = Math.random() * 3 - 1.5;
        this.speedY = Math.random() * 3 - 1.5;
        this.opacity = 1;
        this.idle = idle;

        // 💛 60% жёлтые, 🤍 30% белые, 💜 10% фиолетовые
        const r = Math.random();
        if (r < 0.6) {
            this.color = `hsl(50, 100%, ${Math.random() * 40 + 50}%)`;
            this.shadow = "rgba(255, 255, 150, 0.8)";
        } else if (r < 0.9) {
            this.color = `hsl(60, 20%, ${Math.random() * 70 + 70}%)`;
            this.shadow = "rgba(255, 255, 255, 0.8)";
        } else {
            this.color = `hsl(280, 80%, ${Math.random() * 40 + 60}%)`;
            this.shadow = "rgba(200, 120, 255, 0.8)";
        }
    }

    update() {
        this.x += this.speedX;
        this.y += this.speedY;

        // 🔁 Отражение от краёв экрана
        if (this.x < 0 || this.x > canvas.width) this.speedX *= -1;
        if (this.y < 0 || this.y > canvas.height) this.speedY *= -1;

        this.size *= 0.99;
        this.opacity -= 0.01;
    }

    draw() {
        ctx.beginPath();
        ctx.arc(this.x, this.y, this.size, 0, Math.PI * 2);
        ctx.fillStyle = this.color;
        ctx.shadowColor = this.shadow;
        ctx.shadowBlur = 20;
        ctx.fill();
    }
}

// 🌟 Реакция на движение мыши
window.addEventListener("mousemove", (e) => {
    mouse.x = e.x;
    mouse.y = e.y;
    for (let i = 0; i < 3; i++) {
        particlesArray.push(new Particle(mouse.x, mouse.y));
    }
});

// 🌩️ Постоянное создание искр (вечный эффект ⚡)
setInterval(() => {
    for (let i = 0; i < 2; i++) {
        particlesArray.push(
            new Particle(Math.random() * canvas.width, Math.random() * canvas.height, true)
        );
    }
}, 100); // каждые 100 мс появляются новые искры

function handleParticles() {
    for (let i = 0; i < particlesArray.length; i++) {
        const p = particlesArray[i];
        p.update();
        p.draw();
        if (p.size <= 0.5 || p.opacity <= 0) {
            particlesArray.splice(i, 1);
            i--;
        }
    }
}

function animate() {
    ctx.clearRect(0, 0, canvas.width, canvas.height);
    handleParticles();
    requestAnimationFrame(animate);
}

animate();

window.addEventListener("resize", () => {
    canvas.width = window.innerWidth;
    canvas.height = window.innerHeight;
});
