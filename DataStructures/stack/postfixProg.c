#include <stdio.h>
#include <stdlib.h>
#include <strings.h>
#define MAX 100

char stack[MAX];
char infix[MAX], postfix[MAX];
int top =-1;

void push(char);
char (pop);
int isEmpty();
void inToPost();
int space(char);
int precedence(char);

int main(char *argv[], int argc)
{
	printf("Enter the infix expression: ");
	gets(infix);
	
	inToPost();
	print()
	return 0;
}

void inToPost()
{
	int i, j = 0;
	char symbol, next;
	for(i = 0; i <strlen(infix); i++){
		symbol  = infix[i];
		switch(symbol){
			case '(':
				push(symbol);
				break;
			case ')':
				while((next = pop()) != '(')
					postfix[j++] = next;
			break;
			case '+':
			case '-':
			case '*':
			case '/':
			case '^':
				while(!isEmpty() && precedence(stack[top]) >= precedence(symbol))
					postfix[j++] = pop();
				push(symbol);
				break;
			default:
				postfix[j++] = symbol;			
		}		
	}while(!isEmpty())
		postfix[j++]  = pop();
	postfix[j] = '\0';
	
}

int precedence(char symbol)
{
	switch(symbol)
	{
		//Higher value means greater precedence
		case '^':
			return 3;
		case '/':
		case '*':
			return 2;
		case '+':
		case '-':
			return 1;
		default:
			return 0;
	}
}
int space(char c)
{
	if (c == ' ' || c == '\t')
		return 1;
	else
		return 0;
}

void print()
{
	int i = 0;
	printf("The equivalent postfix expression is: ");
	while (postfix[i]){
		printf("%c", postfix[i++]);
		
	}
	printf('\n');
}
