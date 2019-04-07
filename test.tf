
/*
data "keyring_secret" "test" {
  service = "${keyring_secret.rsaprivkey.service}"
  username = "${keyring_secret.rsaprivkey.username}"
}

output "test_data" {
  value = "${data.keyring_secret.test.value}"
}
*/

resource "keyring_secret" "test" {
  secret = "foobarbazbang!"
}

resource "keyring_secret" "rsaprivkey" {
  service = "terraformrsaprivkey"
  secret = <<HEREDOC
-----BEGIN RSA PRIVATE KEY-----
MIIEpQIBAAKCAQEA6Pf8RnFI+5YbOpbnSufpDlEqv1zUhfYglsCIcJLDTnq5mH24
6zZOQ3ux8OWkddJImY2rNoGM+NOlPthD4YorwH8gq9y/7pwZop4mMzL5e+deHmgo
z6Oi2thGyipKg3sasrJUrAMHgxaTzwi23C4hXakuROJrP1cmRlO9ixR/31fHuxXV
T4jSSMqZYsPV/yFpcKp4kGlo91HcYS3TrWokPKndwhwcNcMvGex1k5sG1oS5cQF0
ZdCA4dl+DVJ/c+4JbOuA9M5eODDbagR9vQLOZBtz2dZD6exWvIGE8pcKeOWodjFS
Y9U1AZtOUS3zVivwmH5haS2zK16ao0vPXnlYHQIDAQABAoIBAD4l4YO9NtX+vjbM
CNmfsjlih7/S+NLfzOPKatt1G2r1LEu0xj7hFtHDzr1P1aakuT9VXqZEuyGp/Dwe
n0Usge1xPHr8gasas3ABNnmUEJ/wIdiEutZhddFhvsixrX0TuxSOpC3bBQcAX9/s
GDP2jmoY9sHjkO6txgUe08KlblNa6h2d91noNh7X0rcLRhV0SKFL9xOOTKFjYDhg
TC32yaoEDEkE7Y+CQ2h83kO5Pitf5ggLSBbVGroD91OcyZ3Lsyz4hrSW2v5r+kj5
dIg0DHDT7+NfEEcvcIZ3kNQH9OZc2QOhllo79btbzc5RuuF5Eb5fB1InLs4+t1+o
7o339jUCgYEA9bQPxeke+aA1kOrG0AXLDsT1/T3KEzvjjQ7v6GC9huEwPBZSZbz2
7I1/Imu/rPKDUPGY3DitjWJ8Z5VG42PtmapjopuJ6FYpe0IgrLliqQXarUTdu6Zm
5bF7pCSi2kv8akD6gahBkrxgpjpdrxUtPKK70x+29UiGPaOvDCMae+sCgYEA8rtN
+K0hQvbkhUks7sTlZf7pSvbwOA5HnVs5l2R4I+1wUHYLtsmg8qG+L5vkwN1IkPTi
FXGaeNzChy5CNZGsdeZo8pzV8VjPxD1yqgaAb1SrzdcZMNt1g7MIdzEh9mtpBnDr
K8FEprAZW0GuB2F873J9m3367EDq4Jo4ng4pIhcCgYEAk/Mr+Elq3HeLKcslkkea
wTb6xNX7I1/RorPW5H+0QE7DA7uRPC/wI/sEDg8BqTuStR/1os4CO3uNW7Y4rF4U
yY4fVt//x8ZyTgVAtaEf9pVO2zz3o7IMjc00nV9uLLLKuJSOA5r0eU4ziITmurGu
vq9jGDslqUdVhjJVCqLspKcCgYEA7ystrczxjCj3jHlIw8/QdQqSVqIwmhs/50We
3pRJaKUpWK6n5iiP5OMIIaMK6BphySrxpc11NofXSRmEdYIfG2C3oe51Q8SNm8As
Pmn/nTMhwcWi9agYB3ed7MvFuielqYoTvt5FRbUYlvmFf5JIZys5pr+gQa/JfxuW
k+5CC2sCgYEA13MS4Iez7ruhbKVoiYJV7YtzkL3/Eqp5HCPgYuw+RfbITGZWj8GM
6WGNQW4Wd0G//3G5p9XtanA2cOf8G2G/f3IPfqhx0Q0yHgn/vI0dJGeD4YCZMruD
qs0juiEmjLQlHnklzmluTJd5dCbIaV1S+qjwojHq+zUnB6Z4/0lmhXk=
-----END RSA PRIVATE KEY-----
HEREDOC
}
