const letters =
  "aábcdeéfghiíjklmnñoópqrstuúüvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ.,;:!?¿¡-";

// Obtener la palabra del primer argumento
const word = process.argv[2] || "";

const startTime = Date.now();

let guess = "";
let attempts = 0;

// Función para obtener un índice aleatorio
const getRandomIndex = () => Math.floor(Math.random() * letters.length);

while (true) {
  // Si la longitud de `guess` excede la de `word`, truncarla
  if (attempts >= word.length) guess = guess.slice(-word.length);

  // Generar una nueva letra aleatoria y añadirla
  guess += letters[getRandomIndex()];
  attempts++;

  // Verificar si ya tenemos una coincidencia completa
  if (guess.length >= word.length && guess.slice(-word.length) === word) break;
}

const endTime = Date.now();
const elapsedTime = (endTime - startTime) / 1000; // Convertir a segundos

console.log("Attempts: " + attempts);
console.log("Time: " + elapsedTime.toFixed(2) + "s");
