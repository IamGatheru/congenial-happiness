struct node *sort_list(struct node *start)
{
	struct node *ptr1, *ptr2;
	int temp;
	ptr1 = start;
	while(ptr1 -> next != NULL)
	{
		ptr2 = ptr1 -> next;
		while(ptr2 != NULL)
		{
			if(ptr1 -> data > ptr2 -> data)
			{
				temp = ptr -> data;
				ptr1-> data = ptr2->data;
				ptr2 -> data = temp;
			}
			ptr2 = ptr2 -> next;
		}
		ptr1 = ptr1 -> next;
	}
	return start;
}
