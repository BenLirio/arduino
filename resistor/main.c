#include <stdio.h>
#include <ctype.h>
#include <string.h>
#define NUM_BANDS 3
enum Band {
	BLACK,
	BROWN,
	RED,
	ORANGE,
	YELLOW,
	GREEN,
	BLUE,
	VIOLET,
	GREY,
	WHITE,
	GOLD,
	SILVER,
	NONE
};
typedef enum Band Band;

void toLowerString(char* s) {
	for (int i = 0; s[i]; i++) {
		s[i] = tolower(s[i]);
	}
}

int getBandVal(char* s) {
	if (strcmp(s, "black") == 0)
		return BLACK;
	if (strcmp(s, "brown") == 0)
		return BROWN;
	if (strcmp(s, "red") == 0)
		return RED;
	if (strcmp(s, "orange") == 0)
		return ORANGE;
	if (strcmp(s, "yellow") == 0)
		return YELLOW;
	if (strcmp(s, "green") == 0)
		return GREEN;
	if (strcmp(s, "blue") == 0)
		return BLUE;
	if (strcmp(s, "violet") == 0)
		return VIOLET;
	if (strcmp(s, "grey") == 0)
		return GREY;
	if (strcmp(s, "white") == 0)
		return WHITE;
	return 0;
}

int main() {
	char bands[NUM_BANDS][128]; 
	int bandsVal[NUM_BANDS];
	for (int i = 0; i < NUM_BANDS; i++) {
		printf("Band %d: ", i+1);
		scanf("%s", bands[i]);
		toLowerString(bands[i]);
		bandsVal[i] = getBandVal(bands[i]);
	}
	int value = (bandsVal[0]*10 + bandsVal[1]);
	for (int i = 0; i < bandsVal[2]; i++) {
		value *= 10;
	}
	printf("\n\n");
	printf("%d \u03A9\n", value);
	float kiloOhms = (float) value/1000;
	printf("%.2f k\u03A9\n", kiloOhms);

	return 0;
}






