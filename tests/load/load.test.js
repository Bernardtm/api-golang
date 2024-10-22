import http from 'k6/http';
import { check, sleep } from 'k6';

// Configuração de opções do teste
export let options = {
  stages: [
    { duration: '30s', target: 5000 }, // Aumenta para 50 VUs em 30 segundos
    { duration: '30s', target: 10000 },  // Mantém 50 VUs por 1 minuto
    { duration: '30s', target: 0 },  // Reduz para 0 VUs em 30 segundos
  ],
};

export default function () {
  // Envia uma requisição GET para o endpoint de healthcheck
  let res = http.get('http://localhost:8080/');

  // Verifica se a resposta tem o código de status 200
  check(res, {
    'status is 200': (r) => r.status === 200,
  });

  // Pausa por um curto período entre as requisições
  sleep(1);
}
