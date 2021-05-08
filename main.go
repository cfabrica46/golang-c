package main

//#include<stdio.h>
//#define pi 3.14159
//int test1() {
//	printf("\n");
//	printf("Test 1:\n");
//	printf("I am in C code now!\n");
//	printf("\n");
//}
//
//int variables() {
//	printf("\n");
//	printf("Variables:\n");
//	int n1;
//	int n2;
//	int r;
//	printf("ingrese el primer valor que sea mayor que 5\n");
//  scanf("%d",&n1);
//	if(n1<=5){
//		printf("introduciste un numero invalido\n");
//		n1 = 0;
//	}
//	printf("ingrese el segundo valor\n");
//  scanf("%d",&n2);
//  r = n1 + n2;
//	printf("el resultado de la suma es: %d\n",r);
//	printf("\n");
//}
//
//int constantes(){
//	printf("\n");
//	printf("Constantes:\n");
//	printf("%f\n",pi);
//	printf("%d\n",(int)pi);
//	printf("\n");
//}
//
//int ciclos(){
//	printf("\n");
//	printf("Ciclos:\n");
//	int n;
//	for(n = 1;n <=10; n++){
//		printf("%d\n",n);
//	}
//	printf("\n");
//}
//
//int switchs(){
//	printf("\n");
//	printf("Switchs:\n");
//	char eleccion;
//	printf("Gato o Perro[g/p]:\n");
//	scanf("%s",eleccion);
//	switch(eleccion){
//		case 'g':
//			printf("miau maiu! :3");
//			break;
//		case 'p':
//			printf("Wof Wof! :3");
//			break;
//		default:
//			printf(">:v");
//			break;
//	}
//	printf("\n");
//}
//
//int arreglos(){
//	printf("\n");
//	printf("arreglos:\n");
//	int size;
//	size = 4;
//	int arreglo[4] = {4,5,6,7};
//	for(int i = 0;i < size; i++){
//		printf("%d\n",arreglo[i]);
//	}
//	printf("\n");
//}
import "C"

func main() {
	C.test1()
	C.variables()
	C.constantes()
	C.ciclos()
	C.switchs()
	C.arreglos()
}
