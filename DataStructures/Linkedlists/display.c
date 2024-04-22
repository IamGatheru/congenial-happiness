struct node *display(struct node *start)
{
	struct node *ptr;
	ptr = start;
	while(ptr != NULL){
		printf("\t %d", ptr -> data);
		ptr = ptr -> next;
	}
	return start;
}
