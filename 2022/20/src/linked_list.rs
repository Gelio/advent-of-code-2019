use std::{cell::RefCell, fmt::Debug, rc::Rc};

type ListItemCell<T> = Rc<RefCell<ListItem<T>>>;

pub struct ListItem<T> {
    prev: Option<ListItemCell<T>>,
    next: Option<ListItemCell<T>>,
    pub value: T,
}

impl<T> Debug for ListItem<T>
where
    T: Debug,
{
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        f.debug_struct("ListItem")
            .field("value", &self.value)
            .finish()
    }
}

impl<T> ListItem<T> {
    fn set_neighbors(&mut self, prev: &ListItemCell<T>, next: &ListItemCell<T>) {
        self.prev = Some(Rc::clone(prev));
        self.next = Some(Rc::clone(next));
    }

    fn prev(&self) -> &ListItemCell<T> {
        self.prev.as_ref().expect("Initialized correctly")
    }

    fn next(&self) -> &ListItemCell<T> {
        self.next.as_ref().expect("Initialized correctly")
    }
}

/// A doubly-linked list with items wrapping around
pub struct List<T> {
    head: ListItemCell<T>,
    length: usize,
    initially_ordered_list_items: Vec<ListItemCell<T>>,
}

impl<T> List<T> {
    pub fn new(items: Vec<T>) -> Self {
        assert!(
            items.len() > 1,
            "Only list with at least 2 elements are supported"
        );

        let list_items: Vec<_> = items
            .into_iter()
            .map(|value| {
                Rc::new(RefCell::new(ListItem {
                    value,
                    next: None,
                    prev: None,
                }))
            })
            .collect();

        let first_item = &list_items[0];
        let last_item = list_items.last().expect("The list has at least 2 elements");
        {
            first_item.borrow_mut().set_neighbors(
                last_item,
                list_items.get(1).expect("The list has at least 2 elements"),
            );

            last_item.borrow_mut().set_neighbors(
                list_items
                    .get(list_items.len() - 2)
                    .expect("The list has at least 2 elements"),
                first_item,
            )
        }

        for (index, item) in list_items.iter().enumerate() {
            if index == 0 || index == list_items.len() - 1 {
                continue;
            }

            item.borrow_mut().set_neighbors(
                list_items
                    .get(index - 1)
                    .expect("We are processing only the middle items"),
                list_items
                    .get(index + 1)
                    .expect("We are processing only the middle items"),
            )
        }

        Self {
            head: Rc::clone(first_item),
            length: list_items.len(),
            initially_ordered_list_items: list_items,
        }
    }

    pub fn move_item(&mut self, item: &ListItemCell<T>, offset: isize) {
        let mut offset = offset
            % (
                // NOTE: account for the moved element being taken out
                // of the collection when moving it
                self.length - 1
            ) as isize;
        if offset == 0 {
            return;
        }

        let mut item_borrowed = item.borrow_mut();
        let item_prev = Rc::clone(item_borrowed.prev());
        let item_next = Rc::clone(item_borrowed.next());
        item_prev.borrow_mut().next = Some(Rc::clone(&item_next));
        item_next.borrow_mut().prev = Some(Rc::clone(&item_prev));

        let mut before_destination_item = Rc::clone(&item_prev);

        while offset != 0 {
            if offset > 0 {
                offset -= 1;
                let next = Rc::clone(before_destination_item.borrow().next());
                before_destination_item = next;
            } else {
                offset += 1;
                let prev = Rc::clone(before_destination_item.borrow().prev());
                before_destination_item = prev;
            }
        }

        let after_destination_item = Rc::clone(before_destination_item.borrow().next());
        item_borrowed.set_neighbors(&before_destination_item, &after_destination_item);
        drop(item_borrowed);

        before_destination_item.borrow_mut().next = Some(Rc::clone(&item));
        after_destination_item.borrow_mut().prev = Some(Rc::clone(&item));

        if Rc::ptr_eq(&self.head, &item) {
            self.head = item_next;
        }
    }

    pub fn initially_ordered_list_items(&self) -> &Vec<ListItemCell<T>> {
        &self.initially_ordered_list_items
    }
}

impl<T: std::fmt::Debug> IntoIterator for List<T> {
    type Item = T;

    type IntoIter = ListIterator<T>;

    fn into_iter(self) -> Self::IntoIter {
        let last_item = self
            .head
            .borrow_mut()
            .prev
            .take()
            .expect("There are at least 2 items in the list");
        last_item.borrow_mut().next = None;
        self.head.borrow_mut().prev = None;

        ListIterator {
            next_item: Some(self.head),
            items_left: self.length,
        }
    }
}

pub struct ListIterator<T> {
    next_item: Option<ListItemCell<T>>,
    items_left: usize,
}

impl<T: std::fmt::Debug> Iterator for ListIterator<T> {
    type Item = T;

    fn next(&mut self) -> Option<Self::Item> {
        if self.items_left == 0 {
            return None;
        }

        self.items_left -= 1;
        let current_item = self.next_item.take().expect("Valid number of items");
        self.next_item = current_item.borrow().next.as_ref().map(Rc::clone);
        if let Some(ref next_item) = self.next_item {
            next_item.borrow_mut().prev = None;
        }

        Some(
            Rc::try_unwrap(current_item)
                .expect("There was an unexpected reference to a list item when iterating")
                .into_inner()
                .value,
        )
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn moves_head() {
        let mut list = List::new(vec![1, 2, -3, 3, -2, 0, 4]);
        list.move_item(&Rc::clone(&list.head), 1);
        assert_eq!(list.head.borrow().value, 2);
        assert_eq!(
            list.into_iter().collect::<Vec<_>>(),
            vec![2, 1, -3, 3, -2, 0, 4]
        );
    }

    #[test]
    fn moves_inner_elements() {
        let mut list = List::new(vec![1, -3, 2, 3, -2, 0, 4]);
        let second_item = Rc::clone(list.head.borrow().next());
        list.move_item(&second_item, -3);
        drop(second_item);
        assert_eq!(list.head.borrow().value, 1);
        assert_eq!(
            list.into_iter().collect::<Vec<_>>(),
            vec![1, 2, 3, -2, -3, 0, 4]
        );
    }
}
