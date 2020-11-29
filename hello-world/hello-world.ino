#define AUDIO 2

int freq = 0;


void setup() {
  pinMode(AUDIO, OUTPUT);
}
int tick = 0;
bool state = true;
void loop() {
  tick += 1;
  if (tick % 1024 == 0) {
    state = !state;
    if (state)
      digitalWrite(AUDIO, HIGH);
    else
      digitalWrite(AUDIO, LOW);
  }
}
